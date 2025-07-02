package persistence

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq" // PostgreSQL驱动
	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/qustavo/sqlhooks/v2"
	"go.uber.org/zap"
)

type PostgresContext struct {
	db *sqlx.DB
}

func (dbc *PostgresContext) Conn(fn func(*sqlx.Conn) error) error {
	conn, err := dbc.db.Connx(context.TODO())
	if err != nil {
		return err
	}
	defer conn.Close()

	return fn(conn)
}

func (dbc *PostgresContext) Transaction(fn func(*sqlx.Tx) error, opts *sql.TxOptions) error {
	tx, err := dbc.db.BeginTxx(context.TODO(), opts)
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

	err = fn(tx)

	return err
}

func newPostgresDB(config config.PostgresConfig, logger *zap.Logger) (DBContext, error) {
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

	return &PostgresContext{db: db}, nil
}
