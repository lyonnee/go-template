package log

import (
	"github.com/lyonnee/go-template/infrastructure/config"
	"github.com/lyonnee/go-template/infrastructure/di"
	"go.uber.org/zap"
)

type Logger = zap.Logger

var (
	logger *Logger
)

func init() {
	config := di.Get[config.Config]()
	newLogger, err := newZapLogger(config.Log)
	if err != nil {
		panic(err)
	}

	logger = newLogger
	di.AddSingleton[*Logger](func() (*Logger, error) {
		return logger, nil
	})
}

// func Initialize(logConfig config.LogConfig,
// ) (*zap.Logger, error) {
// 	newLogger, err := newZapLogger(logConfig)
// 	if err != nil {
// 		return nil, err
// 	}
// 	logger = newLogger

// 	return logger, nil
// }

func Debug(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Debug(msg, fields...)
	}
}

func Info(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Info(msg, fields...)
	}
}

func Warn(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Warn(msg, fields...)
	}
}

func Error(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Error(msg, fields...)
	}
}

func Fatal(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Fatal(msg, fields...)
	}
}

func Panic(msg string, fields ...zap.Field) {
	if logger != nil {
		logger.Panic(msg, fields...)
	}
}

func Sync() error {
	if logger != nil {
		return logger.Sync()
	}
	return nil
}
