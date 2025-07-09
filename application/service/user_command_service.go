package service

import (
	"context"

	"github.com/lyonnee/go-template/bootstrap/di"
	"github.com/lyonnee/go-template/domain/entity"
	"github.com/lyonnee/go-template/domain/repository"
	"github.com/lyonnee/go-template/domain/service"
	"github.com/lyonnee/go-template/infrastructure/auth"
	"github.com/lyonnee/go-template/infrastructure/database"
	"go.uber.org/zap"
)

type UserCommandService struct {
	logger    *zap.Logger
	dbContext database.Database

	userRepo repository.UserRepository

	userDomainService *service.UserService
}

// NewUserApplicationService 创建用户应用服务
func NewUserCommandService() (*UserCommandService, error) {
	return &UserCommandService{
		logger:    di.Get[*zap.Logger](),
		dbContext: di.Get[database.Database](),

		userRepo: di.Get[repository.UserRepository](),

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

// SignUp 用户注册
func (s *UserCommandService) SignUp(ctx context.Context, cmd *SignUpCmd) (*SignUpResult, error) {
	s.logger.Info("Starting user registration",
		zap.String("username", cmd.Username),
		zap.String("email", cmd.Email))

	var user *entity.User
	if err := s.dbContext.Conn(ctx, func(ctx context.Context) error {
		newUser, err := s.userDomainService.CreateUser(ctx, cmd.Username, cmd.Password, cmd.Email, cmd.Phone)
		if err != nil {
			return err
		}

		if err := s.userRepo.Create(ctx, newUser); err != nil {
			return err
		}

		return nil
	}); err != nil {
		s.logger.Error("User registration failed", zap.Error(err), zap.String("username", cmd.Username))
		return nil, err
	}

	// 生成token
	accessToken, err := auth.JWTAuth().GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		s.logger.Error("Failed to generate access token for new user", zap.Error(err), zap.Int64("userId", user.ID))
		return nil, err
	}

	refreshToken, err := auth.JWTAuth().GenerateRefreshToken(user.ID, user.Username)
	if err != nil {
		s.logger.Error("Failed to generate refresh token for new user", zap.Error(err), zap.Int64("userId", user.ID))
		return nil, err
	}

	s.logger.Info("User registration completed successfully",
		zap.String("username", cmd.Username),
		zap.Int64("userId", user.ID))

	return &SignUpResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

// UpdateUsernameCmd 更新用户名命令
type UpdateUsernameCmd struct {
	UserID   int64
	Username string
}

type UpdateResult struct {
	Ok bool
}

// UpdateUsername 更新用户名
func (s *UserCommandService) UpdateUsername(ctx context.Context, cmd *UpdateUsernameCmd) (*entity.User, error) {
	s.logger.Debug("UpdateUsername called",
		zap.Int64("userId", cmd.UserID),
		zap.String("newUsername", cmd.Username))

	var user *entity.User
	if err := s.dbContext.Transaction(ctx, nil, func(ctx context.Context) error {
		// 检查用户是否存在
		user, err := s.userRepo.FindById(ctx, cmd.UserID)
		if err != nil {
			return err
		}

		if err := s.userDomainService.UpdateUsername(ctx, user, cmd.Username); err != nil {
			return err
		}

		if err := s.userRepo.UpdateUsername(ctx, user); err != nil {
			return err
		}

		return nil
	}); err != nil {
		s.logger.Error("Transaction failed during username update", zap.Error(err), zap.Int64("userId", cmd.UserID))
		return nil, err
	}

	return user, nil
}
