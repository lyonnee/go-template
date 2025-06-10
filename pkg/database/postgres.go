package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL驱动
)

func initPostgres(driverName string, dataSourceName string) (*sqlx.DB, error) {
	pgdb, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return pgdb, nil
}
