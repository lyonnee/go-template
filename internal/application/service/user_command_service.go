package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lyonnee/go-template/bootstrap/di"
	"github.com/lyonnee/go-template/internal/domain/entity"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/internal/domain/service"
	"github.com/lyonnee/go-template/internal/infrastructure/persistence"
	"go.uber.org/zap"
)

type UserCommandService struct {
	logger    *zap.Logger
	dbContext persistence.DBContext

	userRepo repository.UserRepository
}

// NewUserApplicationService 创建用户应用服务
func NewUserCommandService() (*UserCommandService, error) {
	return &UserCommandService{
		logger:    di.Get[*zap.Logger](),
		dbContext: di.Get[persistence.DBContext](),

		userRepo: di.Get[repository.UserRepository](),
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
	if err := s.dbContext.Transaction(func(tx *sqlx.Tx) error {
		// 检查用户是否存在
		user, err := s.userRepo.FindById(ctx, cmd.UserID)
		if err != nil {
			s.logger.Error("Failed to find user for username update", zap.Error(err), zap.Int64("userId", cmd.UserID))
			return err
		}

		userDomainService := service.NewUserService(s.userRepo, s.logger)
		if err := userDomainService.UpdateUsername(ctx, user, cmd.Username); err != nil {
			s.logger.Error("Username update failed", zap.Error(err), zap.Int64("userId", cmd.UserID))
			return err
		}

		if err := s.userRepo.UpdateUsername(ctx, user); err != nil {
			s.logger.Error("Failed to update username in repository", zap.Error(err), zap.Int64("userId", cmd.UserID))
			return err
		}

		return nil
	}, nil); err != nil {
		s.logger.Error("Transaction failed during username update", zap.Error(err), zap.Int64("userId", cmd.UserID))
		return nil, err
	}

	return user, nil
}
