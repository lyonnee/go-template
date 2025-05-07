package repository

import "github.com/lyonnee/go-template/internal/domain"

type UserRepository interface {
	FindById(userId int64) (*domain.User, error)
}
