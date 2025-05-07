package repositories

import (
	"github.com/lyonnee/go-template/internal/domain"
	"github.com/lyonnee/go-template/internal/domain/repository"
)

type UserRepository struct {
}

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindById(userId int64) (*domain.User, error) {
	return nil, nil
}
