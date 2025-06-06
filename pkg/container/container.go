package container

import "github.com/samber/do/v2"

type Provider[T any] func() (T, error)

func AddSingletonService[T any](provider Provider[T]) error {
	do.Provide(nil, func(i do.Injector) (T, error) {
		return provider()
	})

	return nil
}

func AddSingletonServiceImpl[I, T any](provider Provider[T]) error {
	do.Provide(nil, func(i do.Injector) (T, error) {
		return provider()
	})

	if err := do.As[T, I](nil); err != nil {
		return err
	}

	return nil
}

func AddTransientService[T any](provider Provider[T]) error {
	do.ProvideTransient(nil, func(i do.Injector) (T, error) {
		return provider()
	})

	return nil
}

func AddTransientServiceImpl[I, T any](provider Provider[T]) error {
	do.ProvideTransient(nil, func(i do.Injector) (T, error) {
		return provider()
	})

	if err := do.As[T, I](nil); err != nil {
		return err
	}

	return nil
}

func GetService[T any]() T {
	return do.MustInvoke[T](nil)
}
