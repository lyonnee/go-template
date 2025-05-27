package log

import (
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"go.uber.org/zap"
)

// ZapSugarLogger 基于zap sugar的日志实现
type ZapSugarLogger struct {
	sugar *zap.SugaredLogger
}

// NewZapSugarLogger 创建新的zap sugar日志实例
func NewZapSugarLogger(zapLogger *zap.Logger) log.Logger {
	return &ZapSugarLogger{
		sugar: zapLogger.Sugar(),
	}
}

// Debug 输出调试信息
func (l *ZapSugarLogger) Debug(msg string, args ...interface{}) {
	l.sugar.Debugw(msg, args...)
}

// Info 输出普通信息
func (l *ZapSugarLogger) Info(msg string, args ...interface{}) {
	l.sugar.Infow(msg, args...)
}

// Warn 输出警告信息
func (l *ZapSugarLogger) Warn(msg string, args ...interface{}) {
	l.sugar.Warnw(msg, args...)
}

// Error 输出错误信息
func (l *ZapSugarLogger) Error(msg string, args ...interface{}) {
	l.sugar.Errorw(msg, args...)
}

// Fatal 输出致命错误信息并退出程序
func (l *ZapSugarLogger) Fatal(msg string, args ...interface{}) {
	l.sugar.Fatalw(msg, args...)
}

// Debugf 输出格式化调试信息
func (l *ZapSugarLogger) Debugf(format string, args ...interface{}) {
	l.sugar.Debugf(format, args...)
}

// Infof 输出格式化普通信息
func (l *ZapSugarLogger) Infof(format string, args ...interface{}) {
	l.sugar.Infof(format, args...)
}

// Warnf 输出格式化警告信息
func (l *ZapSugarLogger) Warnf(format string, args ...interface{}) {
	l.sugar.Warnf(format, args...)
}

// Errorf 输出格式化错误信息
func (l *ZapSugarLogger) Errorf(format string, args ...interface{}) {
	l.sugar.Errorf(format, args...)
}

// Fatalf 输出格式化致命错误信息并退出程序
func (l *ZapSugarLogger) Fatalf(format string, args ...interface{}) {
	l.sugar.Fatalf(format, args...)
}

// DebugKV 输出带键值对的调试信息
func (l *ZapSugarLogger) DebugKV(msg string, keysAndValues ...interface{}) {
	l.sugar.Debugw(msg, keysAndValues...)
}

// InfoKV 输出带键值对的普通信息
func (l *ZapSugarLogger) InfoKV(msg string, keysAndValues ...interface{}) {
	l.sugar.Infow(msg, keysAndValues...)
}

// WarnKV 输出带键值对的警告信息
func (l *ZapSugarLogger) WarnKV(msg string, keysAndValues ...interface{}) {
	l.sugar.Warnw(msg, keysAndValues...)
}

// ErrorKV 输出带键值对的错误信息
func (l *ZapSugarLogger) ErrorKV(msg string, keysAndValues ...interface{}) {
	l.sugar.Errorw(msg, keysAndValues...)
}

// FatalKV 输出带键值对的致命错误信息并退出程序
func (l *ZapSugarLogger) FatalKV(msg string, keysAndValues ...interface{}) {
	l.sugar.Fatalw(msg, keysAndValues...)
}

// WithFields 返回带有指定字段的新logger实例
func (l *ZapSugarLogger) WithFields(keysAndValues ...interface{}) log.Logger {
	return &ZapSugarLogger{
		sugar: l.sugar.With(keysAndValues...),
	}
}

// Sync 同步日志缓冲区
func (l *ZapSugarLogger) Sync() error {
	return l.sugar.Sync()
}
