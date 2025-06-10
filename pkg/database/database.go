package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/lyonnee/go-template/internal/infrastructure/persistence"
	"github.com/qustavo/sqlhooks/v2"
)

type DBContext struct {
	db *sqlx.DB
}

func (dbc *DBContext) NewConn(ctx context.Context) (*sqlx.Conn, error) {
	return dbc.db.Connx(ctx)
}

func (dbc *DBContext) NewTx(ctx context.Context) (*sqlx.Tx, error) {
	return dbc.db.BeginTxx(ctx, nil)
}

func (dbc *DBContext) NewTxWith(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	return dbc.db.BeginTxx(ctx, opts)
}

func NewDB(config *config.PersistenceConfig, logger log.Logger) (persistence.DBContext, error) {
	sql.Register(SQL_LOGGER_DRIVER, sqlhooks.Wrap(pq.Driver{}, &LoggerHooks{Logger: logger}))

	pgDb, err := initPostgres(SQL_LOGGER_DRIVER, config.Postgres.DSN)
	if err != nil {
		return nil, err
	}

	db := pgDb
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	return &DBContext{db: db}, nil
}
