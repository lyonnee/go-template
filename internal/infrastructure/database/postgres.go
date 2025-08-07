package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq" // PostgreSQL驱动
	"github.com/lyonnee/go-template/internal/infrastructure/config"
	"github.com/lyonnee/go-template/pkg/log"
	"github.com/qustavo/sqlhooks/v2"
)

func newPostgresDB(config config.PostgresConfig, logger *log.Logger) (*Database, error) {
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

	return &Database{db: db}, nil
}
