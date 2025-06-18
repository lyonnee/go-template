package di

import "github.com/samber/do/v2"

func AddSingletonService[T any](provider Provider[T]) error {
	return AddSingleton(provider)
}

func AddTransientService[T any](provider Provider[T]) error {
	return AddTransient(provider)
}

func GetService[T any]() T {
	return do.MustInvoke[T](nil)
}
