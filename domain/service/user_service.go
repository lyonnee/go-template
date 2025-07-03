package service

import (
	"context"
	"errors"

	"github.com/lyonnee/go-template/domain/entity"
	domainErrors "github.com/lyonnee/go-template/domain/errors"
	"github.com/lyonnee/go-template/domain/repository"
	"go.uber.org/zap"
)

type UserService struct {
	userRepo repository.UserRepository
	logger   *zap.Logger
}

func NewUserService(
	userRepo repository.UserRepository,
	logger *zap.Logger,
) *UserService {
	return &UserService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (s *UserService) CreateUser(ctx context.Context, username, pwd, email, phone string) (*entity.User, error) {
	// 检查用户名是否已存在
	existingUser, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
		s.logger.Error("Failed to check username existence", zap.String("username", username), zap.Error(err))
		return nil, err
	}
	if existingUser != nil {
		s.logger.Warn("Username already exists", zap.String("username", username))
		return nil, domainErrors.ErrUsernameTaken
	}

	// 检查邮箱是否已存在
	if email != "" {
		existingUser, err := s.userRepo.FindByEmail(ctx, email)
		if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
			s.logger.Error("Failed to check email existence", zap.String("email", email), zap.Error(err))
			return nil, err
		}
		if existingUser != nil {
			s.logger.Warn("Email already exists", zap.String("email", email))
			return nil, domainErrors.ErrEmailTaken
		}
	}

	// 检查手机号是否已存在
	if phone != "" {
		existingUser, err := s.userRepo.FindByPhone(ctx, phone)
		if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
			s.logger.Error("Failed to check phone existence", zap.String("phone", phone), zap.Error(err))
			return nil, err
		}
		if existingUser != nil {
			s.logger.Warn("Phone already exists", zap.String("phone", phone))
			return nil, domainErrors.ErrPhoneTaken
		}
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
			zap.Int64("existingUserId", existingUser.ID),
			zap.Int64("requestingUserId", user.ID))
		return domainErrors.ErrUsernameTaken
	}

	s.logger.Debug("Username is available", zap.String("username", newUsername))

	// 更新用户名
	user.Username = newUsername

	return nil
}
