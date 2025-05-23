package persistence

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/pkg/log"
	"github.com/qustavo/sqlhooks/v2"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Initialize(config config.PersistenceConfig) error {
	pgDb, err := initPostgres(config.Postgres)
	if err != nil {
		return err
	}

	db = pgDb
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	sql.Register("sql-logger", sqlhooks.Wrap(&sqlhooks.Driver{}, &LoggerHooks{Logger: log.Logger()}))

	return nil
}

type Executer interface {
	sqlx.ExecerContext
	sqlx.QueryerContext
}

func NewConn(ctx context.Context) (*sqlx.Conn, error) {
	return db.Connx(ctx)
}

func NewTx(ctx context.Context) (*sqlx.Tx, error) {
	return db.BeginTxx(ctx, nil)
}

func NewTxWith(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	return db.BeginTxx(ctx, opts)
}

type LoggerHooks struct {
	*zap.Logger
}

// Before hook will print the query with it's args and return the context with the timestamp
func (hooks *LoggerHooks) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	if hooks == nil || hooks.Logger == nil {
		return ctx, nil
	}

	return context.WithValue(ctx, "sql_begin", time.Now()), nil
}

// After hook will get the timestamp registered on the Before hook and print the elapsed time
func (hooks *LoggerHooks) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	if hooks == nil || hooks.Logger == nil {
		return ctx, nil
	}
	begin := ctx.Value("sql_begin").(time.Time)

	hooks.Logger.Info("SQL executed",
		zap.String("sql", query),
		zap.Any("args", args),
		zap.Duration("sql_const", time.Since(begin)),
	)
	return ctx, nil
}

func (hooks *LoggerHooks) OnError(_ context.Context, err error, query string, args ...interface{}) error {
	if hooks == nil || hooks.Logger == nil {
		return nil
	}

	hooks.Logger.Error("SQL error",
		zap.String("sql", query),
		zap.Any("args", args),
		zap.Error(err),
	)
	return nil
}
