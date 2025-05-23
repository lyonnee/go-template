package service

import (
	"context"

	"github.com/lyonnee/go-template/internal/domain/entity"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/pkg/persistence"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUserById(ctx context.Context, userId int64) (*entity.User, error) {
	conn, err := persistence.NewConn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	userRepoConn := s.userRepo.WithExecuter(conn)
	user, err := userRepoConn.FindById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
