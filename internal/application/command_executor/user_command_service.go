package command_executor

import (
	"context"

	"github.com/lyonnee/go-template/internal/domain"
	"github.com/lyonnee/go-template/internal/domain/entity"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/lyonnee/go-template/pkg/persistence"
)

type UserCommandService struct {
	userRepo repository.UserRepository
	logger   log.Logger
}

// NewUserApplicationService 创建用户应用服务
func NewUserCommandService(userRepo repository.UserRepository, logger log.Logger) *UserCommandService {
	return &UserCommandService{
		userRepo: userRepo,
		logger:   logger,
	}
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
	s.logger.DebugKV("UpdateUsername called",
		"userId", cmd.UserID,
		"newUsername", cmd.Username)

	// 开启事务
	tx, err := persistence.NewTx(ctx)
	if err != nil {
		s.logger.ErrorKV("Failed to start transaction", "error", err, "userId", cmd.UserID)
		return nil, err
	}
	defer tx.Rollback()

	s.logger.DebugKV("Transaction started for username update", "userId", cmd.UserID)

	userRepoWithTx := s.userRepo.WithExecutor(tx)

	// 检查用户是否存在
	user, err := userRepoWithTx.FindById(ctx, cmd.UserID)
	if err != nil {
		s.logger.ErrorKV("Failed to find user for username update", "error", err, "userId", cmd.UserID)
		return nil, err
	}

	s.logger.DebugKV("User found for username update",
		"userId", cmd.UserID,
		"currentUsername", user.Username)

	if err := userRepoWithTx.UpdateUsername(ctx, user); err != nil {
		s.logger.ErrorKV("Failed to update username", "error", err, "userId", cmd.UserID)
		return nil, err
	}

	userDomainService := domain.NewUserDomainService(userRepoWithTx, s.logger)
	if err := userDomainService.UpdateUsername(ctx, user, cmd.Username); err != nil {
		s.logger.ErrorKV("Username update failed", "error", err, "userId", cmd.UserID)
		return nil, err
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		s.logger.ErrorKV("Failed to commit username update transaction", "error", err, "userId", cmd.UserID)
		return nil, err
	}

	return user, nil
}
