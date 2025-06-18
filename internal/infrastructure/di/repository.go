package di

import (
	"github.com/lyonnee/go-template/internal/infrastructure/persistence"
	"github.com/samber/do/v2"
)

type Repository interface {
	SetExecuter(executer persistence.Executor)
}

func AddSingletonRepository[T any](provider Provider[T]) error {
	return AddSingleton(provider)
}

func AddTransientRepository[T any](provider Provider[T]) error {
	return AddTransient(provider)
}

func GetRepository[T Repository](executer persistence.Executor) T {
	repo := do.MustInvoke[T](nil)
	repo.SetExecuter(executer)
	return repo
}
