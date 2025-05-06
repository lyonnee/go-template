package repositories

import "github.com/lyonnee/go-template/internal/domain"

type UserRepository struct {
}

func NewUserRepository() domain.UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindById(userId int64) (*domain.User, error) {
	return nil, nil
}
