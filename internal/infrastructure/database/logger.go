package database

import (
	"bytes"
	"context"
	"time"

	"github.com/lyonnee/go-template/pkg/log"
	"go.uber.org/zap"
)

const (
	SQL_LOGGER_DRIVER = "sql_logger_driver"
)

type LoggerHooks struct {
	Logger *log.Logger
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
		zap.String("sql", removeEscapes(query)),
		zap.Any("args", args),
		zap.String("duration", time.Since(begin).String()),
	)
	return ctx, nil
}

func (hooks *LoggerHooks) OnError(_ context.Context, err error, query string, args ...interface{}) error {
	if hooks == nil || hooks.Logger == nil {
		return nil
	}

	hooks.Logger.Info("SQL error",
		zap.String("sql", removeEscapes(query)),
		zap.Any("args", args),
		zap.Error(err),
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
