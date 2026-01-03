package commands

import (
	"context"
	"errors"

	"github.com/lyonnee/go-template/internal/domain/entity"
	domainErrors "github.com/lyonnee/go-template/internal/domain/errors"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/internal/domain/service"
	"github.com/lyonnee/go-template/internal/infrastructure/auth"
	"github.com/lyonnee/go-template/internal/infrastructure/database"
	"github.com/lyonnee/go-template/pkg/di"
	"github.com/lyonnee/go-template/pkg/log"
	"go.uber.org/zap"
)

type UserCommandService struct {
	logger *log.Logger
	db     *database.Database

	userDomainService *service.UserService
}

func init() {
	di.AddSingleton[*UserCommandService](NewUserCommandService)
}

// NewUserApplicationService 创建用户应用服务
func NewUserCommandService() (*UserCommandService, error) {
	return &UserCommandService{
		logger: di.Get[*log.Logger](),
		db:     di.Get[*database.Database](),

		userDomainService: di.Get[*service.UserService](),
	}, nil
}

// SignUpCmd 注册命令
type SignUpCmd struct {
	Username string
	Password string
	Email    string
	Phone    string
}

// SignUpResult 注册结果
type SignUpResult struct {
	AccessToken  string
	RefreshToken string
	User         *entity.User
}

// Application Service - 负责编排和唯一性检查
func (s *UserCommandService) SignUp(ctx context.Context, cmd *SignUpCmd) (*SignUpResult, error) {
	var user *entity.User
	var accessToken, refreshToken string

	if err := s.db.WithConnection(ctx, func(ctx context.Context) error {
		userRepo := di.GetRepository[repository.UserRepository](ctx)

		// 1. 先检查唯一性（应用层职责）
		existingUser, err := userRepo.CheckUserFieldsExist(ctx, cmd.Username, cmd.Email, cmd.Phone)
		if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
			return err
		}
		if existingUser {
			return errors.New("user with these details already exists")
		}

		// 2. 调用 domain service 创建用户实体（领域层职责）
		newUser, err := entity.NewUser(cmd.Username, cmd.Password, cmd.Email, cmd.Phone)
		if err != nil {
			return err
		}

		// 3. 持久化（应用层职责）
		if err := userRepo.Create(ctx, newUser); err != nil {
			return err
		}

		// 4. 生成 token
		jwtManager := di.Get[*auth.JWTGenerator]()
		accessToken, err = jwtManager.GenerateAccessToken(newUser.ID, newUser.Username)
		if err != nil {
			return err
		}

		refreshToken, err = jwtManager.GenerateRefreshToken(newUser.ID, newUser.Username)
		if err != nil {
			return err
		}

		user = newUser
		return nil
	}); err != nil {
		return nil, err
	}

	return &SignUpResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

// UpdateUsernameCmd 更新用户名命令
type UpdateUsernameCmd struct {
	UserID   uint64
	Username string
}

type UpdateResult struct {
	Ok bool
}

// UpdateUsername 更新用户名
func (s *UserCommandService) UpdateUsername(ctx context.Context, cmd *UpdateUsernameCmd) (*entity.User, error) {
	s.logger.Debug("UpdateUsername called",
		zap.Uint64("userId", cmd.UserID),
		zap.String("newUsername", cmd.Username))

	var user *entity.User
	if err := s.db.WithTransaction(ctx, nil, func(ctx context.Context) error {
		userRepo := di.GetRepository[repository.UserRepository](ctx)
		// 检查用户是否存在
		user, err := userRepo.FindById(ctx, cmd.UserID)
		if err != nil {
			return err
		}

		// 检查新用户名是否已被其他用户使用
		existingUser, err := userRepo.FindByUsername(ctx, cmd.Username)
		if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
			s.logger.Error("Failed to check username availability", zap.Error(err), zap.String("username", cmd.Username))
			return err
		}
		if existingUser != nil && existingUser.ID != user.ID {
			s.logger.Warn("Username already taken",
				zap.String("username", cmd.Username),
				zap.Uint64("existingUserId", existingUser.ID),
				zap.Uint64("requestingUserId", user.ID))
			return domainErrors.ErrUsernameTaken
		}

		if err := user.UpdateUsername(cmd.Username); err != nil {
			return err
		}

		if err := userRepo.UpdateUsername(ctx, user); err != nil {
			return err
		}

		return nil
	}); err != nil {
		s.logger.Error("Transaction failed during username update", zap.Error(err), zap.Uint64("userId", cmd.UserID))
		return nil, err
	}

	return user, nil
}
