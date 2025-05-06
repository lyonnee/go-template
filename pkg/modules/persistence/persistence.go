package persistence

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lyonnee/go-template/config"
)

var db *sqlx.DB

func Initialize(config config.PersistenceConfig) error {
	err := initPostgres(config.Postgres)
	if err != nil {
		return err
	}

	return nil
}

func NewConn(ctx context.Context) (*sqlx.Conn, error) {
	return db.Connx(ctx)
}

func NewTx(ctx context.Context) (*sqlx.Tx, error) {
	return db.BeginTxx(ctx, nil)
}
