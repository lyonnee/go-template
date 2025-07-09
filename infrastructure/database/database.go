package database

import (
	"context"
	"database/sql"

	"github.com/lyonnee/go-template/infrastructure/config"
	"go.uber.org/zap"
)

type Database interface {
	Conn(ctx context.Context, fn func(context.Context) error) error
	Transaction(ctx context.Context, opts *sql.TxOptions, fn func(context.Context) error) error
	Close() error
}

var db Database

func Initialize(config config.DatabaseConfig, logger *zap.Logger) (Database, error) {
	pgsql, err := newPostgresDB(config.Postgres, logger)
	if err != nil {
		return nil, err
	}

	db = pgsql

	return db, nil
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
