package database

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
)

type DBContext interface {
	sqlx.ExecerContext
	sqlx.QueryerContext
}

type dbContextKeyType struct{}

var dbContextKey = dbContextKeyType{}

var ErrDBContextNotSet = errors.New("dbContext not set, use SetDBContext() to set a dbContext")

func SetDBContext(ctx context.Context, dbContext DBContext) context.Context {
	if dbContext == nil {
		return ctx
	}
	return context.WithValue(ctx, dbContextKey, dbContext)
}

func GetDBContext(ctx context.Context) (DBContext, error) {
	dbContext, ok := ctx.Value(dbContextKey).(DBContext)
	if !ok || dbContext == nil {
		return nil, ErrDBContextNotSet
	}
	return dbContext, nil
}
