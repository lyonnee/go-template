package service

import (
	"context"

	"github.com/lyonnee/go-template/internal/application"
	"github.com/lyonnee/go-template/pkg/auth"
)

type AuthService struct {
	application.AuthService
}

func NewAuthService() application.AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(ctx context.Context, cmd *application.LoginCmd) (*application.LoginResult, error) {
	token, err := auth.GenerateAccessToken(1, cmd.Email)
	if err != nil {
		return nil, err
	}

	return &application.LoginResult{
		Token: token,
	}, nil
}

func (s AuthService) RefreshToken(ctx context.Context, cmd *application.RefreshTokenCmd) (*application.RefreshTokenResult, error) {
	claims, err := auth.ValidateToken(cmd.RefreshToken)
	if err != nil {
		return nil, err
	}

	newToken, err := auth.GenerateAccessToken(claims.UserId, claims.AlternativeID)
	if err != nil {
		return nil, err
	}

	return &application.RefreshTokenResult{
		NewAccessToken: newToken,
	}, nil
}
