package cache

import (
	"context"
	"errors"
	"time"
)

type CacheContext interface {
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string, dest any) error
	Delete(ctx context.Context, key string) error
	DeletePattern(ctx context.Context, pattern string) error
	Exists(ctx context.Context, key string) (bool, error)
	SetWithDefault(ctx context.Context, key string, value any) error
	Increment(ctx context.Context, key string) (int64, error)
	Decrement(ctx context.Context, key string) (int64, error)
	SetNX(ctx context.Context, key string, value any, ttl time.Duration) (bool, error)
}

type cacheContextKeyType struct{}

var cacheContextKey = cacheContextKeyType{}

var ErrCacheContextNotSet = errors.New("cacheContext not set, use SetCacheContext() to set a cacheContext")

func SetCacheContext(ctx context.Context, cacheContext CacheContext) context.Context {
	if cacheContext == nil {
		return ctx
	}
	return context.WithValue(ctx, cacheContextKey, cacheContext)
}

func GetCacheContext(ctx context.Context) (CacheContext, error) {
	cacheContext, ok := ctx.Value(cacheContextKey).(CacheContext)
	if !ok || cacheContext == nil {
		return nil, ErrCacheContextNotSet
	}
	return cacheContext, nil
}
