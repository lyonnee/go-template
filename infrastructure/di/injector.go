package di

import "github.com/samber/do/v2"

type Provider[T any] func() (T, error)

func AddSingleton[T any](provider Provider[T]) error {
	do.Provide(nil, func(do.Injector) (T, error) {
		return provider()
	})

	return nil
}

func AddTransient[T any](provider Provider[T]) error {
	do.ProvideTransient(nil, func(do.Injector) (T, error) {
		return provider()
	})

	return nil
}

func Get[T any]() T {
	return do.MustInvoke[T](nil)
}
