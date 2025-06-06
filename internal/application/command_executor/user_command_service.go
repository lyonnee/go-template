package command_executor

import (
	"context"

	"github.com/lyonnee/go-template/internal/domain"
	"github.com/lyonnee/go-template/internal/domain/entity"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/lyonnee/go-template/pkg/container"
	"github.com/lyonnee/go-template/pkg/persistence"
)

type UserCommandService struct {
	userRepo repository.UserRepository
	logger   log.Logger
}

// NewUserApplicationService 创建用户应用服务
func NewUserCommandService() (*UserCommandService, error) {
	return &UserCommandService{
		userRepo: container.GetService[repository.UserRepository](),
		logger:   container.GetService[log.Logger](),
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
	// 使用事务执行器
	userRepoWithTx := s.userRepo.WithExecutor(tx)

	// 检查用户是否存在
	user, err := userRepoWithTx.FindById(ctx, cmd.UserID)
	if err != nil {
		s.logger.ErrorKV("Failed to find user for username update", "error", err, "userId", cmd.UserID)
		return nil, err
	}

	userDomainService := domain.NewUserDomainService(userRepoWithTx, s.logger)
	if err := userDomainService.UpdateUsername(ctx, user, cmd.Username); err != nil {
		s.logger.ErrorKV("Username update failed", "error", err, "userId", cmd.UserID)
		return nil, err
	}

	if err := userRepoWithTx.UpdateUsername(ctx, user); err != nil {
		s.logger.ErrorKV("Failed to update username in repository", "error", err, "userId", cmd.UserID)
		return nil, err
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		s.logger.ErrorKV("Failed to commit username update transaction", "error", err, "userId", cmd.UserID)
		return nil, err
	}

	return user, nil
}
