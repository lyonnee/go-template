package database

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type dbContextKeyType struct{}

var dbContextKey = dbContextKeyType{}

var ErrDBContextNotSet = errors.New("dbContext not set, use SetDBContext() to set a dbContext")

type DBContext interface {
	WithContext(ctx context.Context, fn func(context.Context) error) error
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}

func SetDBContext(ctx context.Context, db *gorm.DB) context.Context {
	if db == nil {
		return ctx
	}
	return context.WithValue(ctx, dbContextKey, db)
}

func GetDBContext(ctx context.Context) (*gorm.DB, error) {
	db, ok := ctx.Value(dbContextKey).(*gorm.DB)
	if !ok || db == nil {
		return nil, ErrDBContextNotSet
	}
	return db, nil
}
