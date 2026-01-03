package di

import (
	"context"

	"github.com/samber/do/v2"
)

type Provider[T any] func() (T, error)

func AddSingleton[T any](provider Provider[T]) error {
	do.Provide(nil, func(do.Injector) (T, error) {
		return provider()
	})

	return nil
}

func AddSingletonImpl[I, T any](provider Provider[T]) error {
	do.Provide(nil, func(do.Injector) (T, error) {
		return provider()
	})

	err := do.As[T, I](nil)
	if err != nil {
		return err
	}

	return nil
}

func AddTransient[T any](provider Provider[T]) error {
	do.ProvideTransient(nil, func(do.Injector) (T, error) {
		return provider()
	})

	return nil
}

func AddTransientImpl[I, T any](provider Provider[T]) error {
	do.ProvideTransient(nil, func(do.Injector) (T, error) {
		return provider()
	})

	return do.As[T, I](nil)
}

func Get[T any]() T {
	return do.MustInvoke[T](nil)
}

type Repository interface {
	SetContext(ctx context.Context)
}

func GetRepository[T Repository](ctx context.Context) T {
	instance := do.MustInvoke[T](nil)
	instance.SetContext(ctx)
	return instance
}
