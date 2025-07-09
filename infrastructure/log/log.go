package log

import (
	"github.com/lyonnee/go-template/infrastructure/config"
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
)

func Initialize(logConfig config.LogConfig,
) (*zap.Logger, error) {
	logger, err := newZapLogger(logConfig)
	if err != nil {
		return nil, err
	}
	Logger = logger

	return Logger, nil
}

func Debug(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Debug(msg, fields...)
	}
}

func Info(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Info(msg, fields...)
	}
}

func Warn(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Warn(msg, fields...)
	}
}

func Error(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Error(msg, fields...)
	}
}

func Fatal(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Fatal(msg, fields...)
	}
}

func Panic(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Panic(msg, fields...)
	}
}

func Sync() error {
	if Logger != nil {
		return Logger.Sync()
	}
	return nil
}
