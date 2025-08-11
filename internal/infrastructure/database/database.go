package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lyonnee/go-template/internal/infrastructure/config"
	"github.com/lyonnee/go-template/pkg/di"
	"github.com/lyonnee/go-template/pkg/log"
)

type Database struct {
	db *sqlx.DB
}

// func (dbc *Database) Conn(ctx context.Context) (DBExecutor, error) {
// 	return dbc.db.Connx(ctx)
// }

// func (dbc *Database) Transaction(ctx context.Context, opts *sql.TxOptions) (DBExecutor, error) {
// 	return dbc.db.BeginTxx(ctx, opts)
// }

func (dbc *Database) Conn(ctx context.Context, fn func(context.Context) error) error {
	conn, err := dbc.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	return fn(SetDBExecutor(ctx, conn))
}

func (dbc *Database) Transaction(ctx context.Context, opts *sql.TxOptions, fn func(context.Context) error) error {
	tx, err := dbc.db.BeginTxx(ctx, opts)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(SetDBExecutor(ctx, tx))

	return err
}

func (dbc *Database) Close() error {
	if dbc.db != nil {
		return dbc.db.Close()
	}
	return nil
}

var db *Database

func init() {
	config := di.Get[config.Config]()
	logger := di.Get[*log.Logger]()

	pgsql, err := newPostgresDB(config.Database.Postgres, logger)
	if err != nil {
		panic("Failed to initialize PostgreSQL database: " + err.Error())
	}

	db = pgsql

	di.AddSingleton[*Database](func() (*Database, error) {
		return db, nil
	})
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
