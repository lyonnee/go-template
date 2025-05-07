package service

import (
	"context"

	"github.com/lyonnee/go-template/internal/application"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(ctx context.Context, cmd *application.LoginCmd) (*application.LoginResult, error) {
	return nil, nil
}
