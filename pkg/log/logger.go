package log

import (
	"github.com/lyonnee/go-template/config"
	"go.uber.org/zap"
)

var logger *zap.Logger

func Logger() *zap.Logger {
	return logger
}

func Initialize(logConfig config.LogConfig) error {
	zapLogger, err := newZap(
		logConfig.EnableToConsole,
		logConfig.ToConsoleLevel,

		logConfig.Filename,
		logConfig.ToFileLevel,
		logConfig.MaxSize,
		logConfig.MaxBackups,
		logConfig.MaxAge,
	)

	if err != nil {
		return err
	}

	logger = zapLogger
	return nil
}

func Sync() {
	logger.Sync()
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
