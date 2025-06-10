package database

import (
	"bytes"
	"context"
	"time"

	"github.com/lyonnee/go-template/internal/infrastructure/log"
)

const (
	SQL_LOGGER_DRIVER = "sql_logger_driver"
)

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
