package persistence

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/qustavo/sqlhooks/v2"
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

	// Note: LoggerHooks will be initialized separately with proper logger injection
	sql.Register("sql-logger", sqlhooks.Wrap(&sqlhooks.Driver{}, &LoggerHooks{}))

	return nil
}

// SetLogger sets the logger for SQL hooks
func SetLogger(logger log.Logger) {
	// Update the registered driver with the new logger
	sql.Register("sql-logger", sqlhooks.Wrap(&sqlhooks.Driver{}, &LoggerHooks{Logger: logger}))
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
	Logger log.Logger
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

	hooks.Logger.InfoKV("SQL executed",
		"sql", query,
		"args", args,
		"duration", time.Since(begin),
	)
	return ctx, nil
}

func (hooks *LoggerHooks) OnError(_ context.Context, err error, query string, args ...interface{}) error {
	if hooks == nil || hooks.Logger == nil {
		return nil
	}

	hooks.Logger.ErrorKV("SQL error",
		"sql", query,
		"args", args,
		"error", err,
	)
	return nil
}
