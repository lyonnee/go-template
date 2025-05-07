package service

import (
	"context"

	"github.com/lyonnee/go-template/internal/domain"
)

type UserService struct {
}

func (s *UserService) NewUser(ctx context.Context, userId int64, username string) (*domain.User, error) {
	return nil, nil
}
