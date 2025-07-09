package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq" // PostgreSQL驱动
	"github.com/lyonnee/go-template/infrastructure/config"
	"github.com/lyonnee/go-template/infrastructure/log"
	"github.com/qustavo/sqlhooks/v2"
	"go.uber.org/zap"
)

type PostgresDB struct {
	db *sqlx.DB
}

func (dbc *PostgresDB) Conn(ctx context.Context, fn func(context.Context) error) error {
	conn, err := dbc.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	return fn(SetDBExecutor(ctx, conn))
}

func (dbc *PostgresDB) Transaction(ctx context.Context, opts *sql.TxOptions, fn func(context.Context) error) error {
	tx, err := dbc.db.BeginTxx(ctx, opts)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			log.Logger.Error("Failed to rollback transaction", zap.Any("err", p))
		} else if err != nil {
			_ = tx.Rollback()
			log.Logger.Error("Failed to rollback transaction", zap.Error(err))
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(SetDBExecutor(ctx, tx))

	return err
}

func (dbc *PostgresDB) Close() error {
	if dbc.db != nil {
		return dbc.db.Close()
	}
	return nil
}

func newPostgresDB(config config.PostgresConfig, logger *zap.Logger) (Database, error) {
	sql.Register(SQL_LOGGER_DRIVER, sqlhooks.Wrap(pq.Driver{}, &LoggerHooks{Logger: logger}))

	pgDb, err := sqlx.Connect(SQL_LOGGER_DRIVER, config.DSN)
	if err != nil {
		return nil, err
	}

	db := pgDb
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	return &PostgresDB{db: db}, nil
}
