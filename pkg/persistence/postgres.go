package persistence

import (
	"github.com/jmoiron/sqlx"
	"github.com/lyonnee/go-template/config"
)

func initPostgres(config config.PostgresConfig) error {
	pgdb, err := sqlx.Connect("postgres", config.DSN)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(5)

	db = pgdb

	return nil
}
