package database

import (
	"context"
	"errors"
	"time"

	"github.com/lyonnee/go-template/pkg/log"
	"go.uber.org/zap"
	gormlogger "gorm.io/gorm/logger"
)

type GormLogger struct {
	Logger   *log.Logger
	LogLevel gormlogger.LogLevel
}

func NewGormLogger(logger *log.Logger) *GormLogger {
	return &GormLogger{
		Logger:   logger,
		LogLevel: gormlogger.Info,
	}
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Info {
		l.Logger.Sugar().Infof(msg, data...)
	}
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Warn {
		l.Logger.Sugar().Warnf(msg, data...)
	}
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Error {
		l.Logger.Sugar().Errorf(msg, data...)
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil && !errors.Is(err, gormlogger.ErrRecordNotFound) {
		l.Logger.Error("SQL error",
			zap.String("sql", sql),
			zap.Int64("rows", rows),
			zap.String("duration", elapsed.String()),
			zap.Error(err),
		)
		return
	}

	l.Logger.Info("SQL executed",
		zap.String("sql", sql),
		zap.Int64("rows", rows),
		zap.String("duration", elapsed.String()),
	)
}
