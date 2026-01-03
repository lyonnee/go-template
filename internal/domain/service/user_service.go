package service

import (
	"github.com/lyonnee/go-template/internal/domain/entity"
	"github.com/lyonnee/go-template/pkg/di"
	"github.com/lyonnee/go-template/pkg/log"
	"go.uber.org/zap"
)

type UserService struct {
	logger *log.Logger
}

func init() {
	di.AddSingleton[*UserService](NewUserService)
}

func NewUserService() (*UserService, error) {
	return &UserService{
		logger: di.Get[*log.Logger](),
	}, nil
}

func (s *UserService) NewUser(username, pwd, email, phone string) (*entity.User, error) {
	user, err := entity.NewUser(username, pwd, email, phone)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUsername(user *entity.User, newUsername string) error {
	s.logger.Debug("Username is available", zap.String("username", newUsername))

	// 更新用户名
	user.Username = newUsername

	return nil
}
