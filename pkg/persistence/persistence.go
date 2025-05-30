package persistence

import (
	"bytes"
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/qustavo/sqlhooks/v2"
)

var db *sqlx.DB

func Initialize(config config.PersistenceConfig, logger log.Logger) error {
	sql.Register(SQL_LOGGER_DRIVER, sqlhooks.Wrap(pq.Driver{}, &LoggerHooks{Logger: logger}))

	pgDb, err := initPostgres(config.Postgres)
	if err != nil {
		return err
	}

	db = pgDb
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	return nil
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
		"sql", removeEscapes(query),
		"args", args,
		"duration", time.Since(begin).String(),
	)
	return ctx, nil
}

func (hooks *LoggerHooks) OnError(_ context.Context, err error, query string, args ...interface{}) error {
	if hooks == nil || hooks.Logger == nil {
		return nil
	}

	hooks.Logger.ErrorKV("SQL error",
		"sql", removeEscapes(query),
		"args", args,
		"error", err,
	)
	return nil
}

func removeEscapes(s string) string {
	var buf bytes.Buffer
	for _, r := range s {
		switch r {
		case '\n', '\t', '\\':
			continue // 跳过转义字符
		default:
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
