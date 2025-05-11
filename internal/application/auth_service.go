package application

import "context"

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

type AuthService interface {
	Login(ctx context.Context, cmd *LoginCmd) (*LoginResult, error)
	RefreshToken(ctx context.Context, cmd *RefreshTokenCmd) (*RefreshTokenResult, error)
}
