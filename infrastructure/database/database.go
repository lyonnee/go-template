package database

import (
	"context"
	"database/sql"

	"github.com/lyonnee/go-template/infrastructure/config"
	"github.com/lyonnee/go-template/infrastructure/di"
	"github.com/lyonnee/go-template/infrastructure/log"
)

type Database interface {
	Conn(ctx context.Context, fn func(context.Context) error) error
	Transaction(ctx context.Context, opts *sql.TxOptions, fn func(context.Context) error) error
	Close() error
}

var db Database

func init() {
	config := di.Get[config.Config]()
	logger := di.Get[*log.Logger]()

	pgsql, err := newPostgresDB(config.Database.Postgres, logger)
	if err != nil {
		panic(err)
	}

	db = pgsql

	di.AddSingleton[Database](func() (Database, error) {
		return db, nil
	})
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
