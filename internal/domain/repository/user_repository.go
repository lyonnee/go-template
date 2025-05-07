package repository

import (
	"context"

	"github.com/lyonnee/go-template/internal/domain"
)

type UserRepository interface {
	FindById(ctx context.Context, userId int64) (*domain.User, error)
}
