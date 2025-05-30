package query_executor

import (
	"context"

	"github.com/lyonnee/go-template/internal/infrastructure/log"

	"github.com/lyonnee/go-template/internal/domain/entity"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/pkg/persistence"
)

// UserApplicationService 用户应用服务
type UserQueryService struct {
	userRepo repository.UserRepository
	logger   log.Logger
}

// NewUserApplicationService 创建用户应用服务
func NewUserQueryService(userRepo repository.UserRepository, logger log.Logger) *UserQueryService {
	return &UserQueryService{
		userRepo: userRepo,
		logger:   logger,
	}
}

// GetUserById 根据ID获取用户
func (s *UserQueryService) GetUserById(ctx context.Context, userId int64) (*entity.User, error) {
	s.logger.DebugKV("GetUserById called", "userId", userId)

	conn, err := persistence.NewConn(ctx)
	if err != nil {
		s.logger.ErrorKV("Failed to create database connection", "error", err, "userId", userId)
		return nil, err
	}
	defer conn.Close()

	userRepoConn := s.userRepo.WithExecutor(conn)
	user, err := userRepoConn.FindById(ctx, userId)
	if err != nil {
		s.logger.ErrorKV("Failed to find user by ID", "error", err, "userId", userId)
		return nil, err
	}

	s.logger.InfoKV("User found successfully", "userId", userId, "username", user.Username)
	return user, nil
}
