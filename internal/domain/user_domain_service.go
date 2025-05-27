package domain

import (
	"context"
	"errors"

	"github.com/lyonnee/go-template/internal/domain/entity"
	domainErrors "github.com/lyonnee/go-template/internal/domain/errors"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
)

type UserDomainService struct {
	userRepo repository.UserRepository
	logger   log.Logger
}

func NewUserDomainService(
	userRepo repository.UserRepository,
	logger log.Logger,
) *UserDomainService {
	return &UserDomainService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (s *UserDomainService) CreateUser(ctx context.Context, username, pwd, email, phone string) (*entity.User, error) {
	// 检查用户名是否已存在
	existingUser, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
		s.logger.ErrorKV("Failed to check username existence",
			"username", username,
			"error", err)
		return nil, err
	}
	if existingUser != nil {
		s.logger.WarnKV("Username already exists", "username", username)
		return nil, domainErrors.ErrUsernameTaken
	}

	// 检查邮箱是否已存在
	if email != "" {
		existingUser, err := s.userRepo.FindByEmail(ctx, email)
		if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
			s.logger.ErrorKV("Failed to check email existence",
				"email", email,
				"error", err)
			return nil, err
		}
		if existingUser != nil {
			s.logger.WarnKV("Email already exists", "email", email)
			return nil, domainErrors.ErrEmailTaken
		}
	}

	// 检查手机号是否已存在
	if phone != "" {
		existingUser, err := s.userRepo.FindByPhone(ctx, phone)
		if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
			s.logger.ErrorKV("Failed to check phone existence",
				"phone", phone,
				"error", err)
			return nil, err
		}
		if existingUser != nil {
			s.logger.WarnKV("Phone already exists", "phone", phone)
			return nil, domainErrors.ErrPhoneTaken
		}
	}

	user, err := entity.NewUser(username, pwd, email, phone)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserDomainService) UpdateUsername(ctx context.Context, user *entity.User, newUsername string) error {
	// 检查新用户名是否已被其他用户使用
	existingUser, err := s.userRepo.FindByUsername(ctx, newUsername)
	if err != nil && !errors.Is(err, domainErrors.ErrUserNotFound) {
		s.logger.ErrorKV("Failed to check username availability", "error", err, "username", newUsername)
		return err
	}
	if existingUser != nil && existingUser.ID != user.ID {
		s.logger.WarnKV("Username already taken",
			"username", newUsername,
			"existingUserId", existingUser.ID,
			"requestingUserId", user.ID)
		return domainErrors.ErrUsernameTaken
	}

	s.logger.DebugKV("Username is available", "username", newUsername)

	// 更新用户名
	user.Username = newUsername

	return nil
}
