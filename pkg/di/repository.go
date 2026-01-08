package di

import (
	"context"

	"github.com/samber/do/v2"
)

type Repository interface {
	SetContext(ctx context.Context)
}

func GetRepository[T Repository](ctx context.Context) T {
	instance := do.MustInvoke[T](nil)
	instance.SetContext(ctx)
	return instance
}
