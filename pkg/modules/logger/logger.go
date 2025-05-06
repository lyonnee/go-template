package logger

import (
	"github.com/lyonnee/go-template/config"
	"go.uber.org/zap"
)

func Initialize(logConfig config.LogConfig) {
	zapLogger = newZap(
		logConfig.Filename,
		logConfig.Level,
		logConfig.MaxSize,
		logConfig.MaxBackups,
		logConfig.MaxAge,
	)
}

func Sync() {
	syncZap()
}

func Debug(msg string, fields ...zap.Field) {
	zapLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	zapLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zapLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zapLogger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	zapLogger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	zapLogger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	zapLogger.Fatal(msg, fields...)
}
