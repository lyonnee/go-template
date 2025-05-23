package application

import (
	"context"

	"github.com/lyonnee/go-template/pkg/auth"
)

type LoginCmd struct {
	PhoneNumber string
	Email       string
	Password    string
}

type LoginResult struct {
	Token string
}

type RefreshTokenCmd struct {
	RefreshToken string
}

type RefreshTokenResult struct {
	NewAccessToken string
}

func Login(ctx context.Context, cmd *LoginCmd) (*LoginResult, error) {
	token, err := auth.GenerateAccessToken(1, cmd.Email)
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		Token: token,
	}, nil
}

func RefreshToken(ctx context.Context, cmd *RefreshTokenCmd) (*RefreshTokenResult, error) {
	claims, err := auth.ValidateToken(cmd.RefreshToken)
	if err != nil {
		return nil, err
	}

	newToken, err := auth.GenerateAccessToken(claims.UserId, claims.AlternativeID)
	if err != nil {
		return nil, err
	}

	return &RefreshTokenResult{
		NewAccessToken: newToken,
	}, nil
}
