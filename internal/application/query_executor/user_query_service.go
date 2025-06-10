package query_executor

import (
	"context"

	"github.com/lyonnee/go-template/internal/infrastructure/di"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/lyonnee/go-template/internal/infrastructure/persistence"

	"github.com/lyonnee/go-template/internal/domain/entity"
	"github.com/lyonnee/go-template/internal/domain/repository"
)

// UserApplicationService 用户应用服务
type UserQueryService struct {
	logger    log.Logger
	dbContext persistence.DBContext

	userRepo repository.UserRepository
}

// NewUserApplicationService 创建用户应用服务
func NewUserQueryService() (*UserQueryService, error) {
	return &UserQueryService{
		logger:    di.GetService[log.Logger](),
		dbContext: di.GetService[persistence.DBContext](),

		userRepo: di.GetService[repository.UserRepository](),
	}, nil
}

// GetUserById 根据ID获取用户
func (s *UserQueryService) GetUserById(ctx context.Context, userId int64) (*entity.User, error) {
	s.logger.DebugKV("GetUserById called", "userId", userId)

	conn, err := s.dbContext.NewConn(ctx)
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
