package service

import (
	"context"
	"errors"

	"github.com/lyonnee/go-template/internal/domain/entity"
	domainErrors "github.com/lyonnee/go-template/internal/domain/errors"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/pkg/di"
	"github.com/lyonnee/go-template/pkg/log"
	"go.uber.org/zap"
)

type UserService struct {
	logger   *log.Logger
	userRepo repository.UserRepository
}

func init() {
	di.AddSingleton[*UserService](NewUserService)
}

func NewUserService() (*UserService, error) {
	return &UserService{
		logger:   di.Get[*log.Logger](),
		userRepo: di.Get[repository.UserRepository](),
	}, nil
}

func (s *UserService) NewUser(ctx context.Context, username, pwd, email, phone string) (*entity.User, error) {
	// 检查用户名是否已存在
	existingUser, err := s.userRepo.CheckUserFieldsExist(ctx, username, email, phone)
	if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
		s.logger.Error("Failed to check user fields existence", zap.String("username", username), zap.String("email", email), zap.String("phone", phone), zap.Error(err))
		return nil, err
	}

	if existingUser {
		s.logger.Warn("User with these details already exists", zap.String("username", username), zap.String("email", email), zap.String("phone", phone))
		return nil, errors.New("user with these details already exists")
	}

	user, err := entity.NewUser(username, pwd, email, phone)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUsername(ctx context.Context, user *entity.User, newUsername string) error {
	// 检查新用户名是否已被其他用户使用
	existingUser, err := s.userRepo.FindByUsername(ctx, newUsername)
	if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
		s.logger.Error("Failed to check username availability", zap.Error(err), zap.String("username", newUsername))
		return err
	}
	if existingUser != nil && existingUser.ID != user.ID {
		s.logger.Warn("Username already taken",
			zap.String("username", newUsername),
			zap.Uint64("existingUserId", existingUser.ID),
			zap.Uint64("requestingUserId", user.ID))
		return domainErrors.ErrUsernameTaken
	}

	s.logger.Debug("Username is available", zap.String("username", newUsername))

	// 更新用户名
	user.Username = newUsername

	return nil
}
