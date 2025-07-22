package database

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
)

const DBExecutorKey = "db_executor"

var ErrDBExecutorNotSet = errors.New("dbExecutor not set, use SetDBExecutor() to set a dbExecutor")

type DBExecutor interface {
	sqlx.ExecerContext
	sqlx.QueryerContext
}

func SetDBExecutor(ctx context.Context, dbExecutor DBExecutor) context.Context {
	if dbExecutor == nil {
		return ctx
	}
	return context.WithValue(ctx, DBExecutorKey, dbExecutor)
}

func GetDBExecutor(ctx context.Context) (DBExecutor, error) {
	dbExecutor, ok := ctx.Value(DBExecutorKey).(DBExecutor)
	if !ok || dbExecutor == nil {
		return nil, ErrDBExecutorNotSet
	}
	return dbExecutor, nil
}
