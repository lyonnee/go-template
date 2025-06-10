package persistence

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DBContext interface {
	NewConn(ctx context.Context) (*sqlx.Conn, error)
	NewTx(ctx context.Context) (*sqlx.Tx, error)
	NewTxWith(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}
