package persistence

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lyonnee/go-template/config"
	"go.uber.org/zap"
)

type DBContext interface {
	Conn(fn func(*sqlx.Conn) error) error
	Transaction(fn func(*sqlx.Tx) error, opts *sql.TxOptions) error
}

type Executor interface {
	sqlx.ExecerContext
	sqlx.QueryerContext
}

func Initialize(config config.PersistenceConfig, logger *zap.Logger) (DBContext, error) {
	db, err := newPostgresDB(config.Postgres, logger)
	if err != nil {
		return nil, err
	}

	return db, nil
}
