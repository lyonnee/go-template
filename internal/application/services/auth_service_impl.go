package services

import "github.com/lyonnee/go-template/internal/application"

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(cmd *application.LoginCmd) error {
	return nil
}
