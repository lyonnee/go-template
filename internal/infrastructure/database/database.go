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

func (dbc *Database) WithConnection(ctx context.Context, fn func(context.Context) error) error {
	conn, err := dbc.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	return fn(SetDBContext(ctx, conn))
}

func (dbc *Database) WithTransaction(ctx context.Context, opts *sql.TxOptions, fn func(context.Context) error) error {
	tx, err := dbc.db.BeginTxx(ctx, opts)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(SetDBContext(ctx, tx))

	return err
}

func (dbc *Database) Close() error {
	if dbc.db != nil {
		return dbc.db.Close()
	}
	return nil
}

// CloseWithContext attempts to close the underlying DB while respecting the given context's deadline.
// Since sqlx.DB.Close() itself doesn't accept a context and usually returns quickly,
// we guard it with a goroutine and select on context timeout to avoid blocking shutdown.
func (dbc *Database) CloseWithContext(ctx context.Context) error {
	if dbc.db == nil {
		return nil
	}
	done := make(chan error, 1)
	go func() {
		done <- dbc.db.Close()
	}()
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
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

// CloseWithContext closes the global DB instance with context awareness.
func CloseWithContext(ctx context.Context) error {
	if db != nil {
		return db.CloseWithContext(ctx)
	}
	return nil
}
