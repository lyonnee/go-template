package persistence

import "github.com/jmoiron/sqlx"

type Executor interface {
	sqlx.ExecerContext
	sqlx.QueryerContext
}
