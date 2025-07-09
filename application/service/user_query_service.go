package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/lyonnee/go-template/bootstrap/di"
	"github.com/lyonnee/go-template/domain/entity"
	"github.com/lyonnee/go-template/domain/repository"
	"github.com/lyonnee/go-template/infrastructure/database"
)

// UserApplicationService 用户应用服务
type UserQueryService struct {
	logger    *zap.Logger
	dbContext database.Database

	userRepo repository.UserRepository
}

// NewUserApplicationService 创建用户应用服务
func NewUserQueryService() (*UserQueryService, error) {
	return &UserQueryService{
		logger:    di.Get[*zap.Logger](),
		dbContext: di.Get[database.Database](),

		userRepo: di.Get[repository.UserRepository](),
	}, nil
}

// GetUserById 根据ID获取用户
func (s *UserQueryService) GetUserById(ctx context.Context, userId int64) (*entity.User, error) {
	s.logger.Debug("GetUserById called", zap.Int64("userId", userId))

	var user *entity.User
	if err := s.dbContext.Conn(ctx, func(ctx context.Context) error {
		userInfo, err := s.userRepo.FindById(ctx, userId)
		if err != nil {
			return err
		}

		user = userInfo

		return nil
	}); err != nil {
		s.logger.Error("Database connection failed", zap.Error(err), zap.Int64("userId", userId))
		return nil, err
	}

	s.logger.Info("User found successfully", zap.Int64("userId", userId), zap.String("username", user.Username))
	return user, nil
}
