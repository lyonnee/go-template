package di

import (
	"github.com/samber/do/v2"
)

func AddSingletonController[T any](provider Provider[T]) error {
	return AddSingleton(provider)
}

func AddTransientController[T any](provider Provider[T]) error {
	return AddTransient(provider)
}

func GetController[T any]() T {
	return do.MustInvoke[T](nil)
}
