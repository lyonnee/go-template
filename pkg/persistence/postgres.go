package persistence

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL驱动
	"github.com/lyonnee/go-template/config"
)

const (
	SQL_LOGGER_DRIVER = "sql_logger_driver"
)

func initPostgres(config config.PostgresConfig) (*sqlx.DB, error) {
	pgdb, err := sqlx.Connect(SQL_LOGGER_DRIVER, config.DSN)
	if err != nil {
		return nil, err
	}

	return pgdb, nil
}
