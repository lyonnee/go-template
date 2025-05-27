package log

// NoopLogger 空操作日志实现，用于测试
type NoopLogger struct{}

// NewNoopLogger 创建新的空操作日志实例
func NewNoopLogger() Logger {
	return &NoopLogger{}
}

// Debug 空操作
func (l *NoopLogger) Debug(msg string, args ...interface{}) {}

// Info 空操作
func (l *NoopLogger) Info(msg string, args ...interface{}) {}

// Warn 空操作
func (l *NoopLogger) Warn(msg string, args ...interface{}) {}

// Error 空操作
func (l *NoopLogger) Error(msg string, args ...interface{}) {}

// Fatal 空操作
func (l *NoopLogger) Fatal(msg string, args ...interface{}) {}

// Debugf 空操作
func (l *NoopLogger) Debugf(format string, args ...interface{}) {}

// Infof 空操作
func (l *NoopLogger) Infof(format string, args ...interface{}) {}

// Warnf 空操作
func (l *NoopLogger) Warnf(format string, args ...interface{}) {}

// Errorf 空操作
func (l *NoopLogger) Errorf(format string, args ...interface{}) {}

// Fatalf 空操作
func (l *NoopLogger) Fatalf(format string, args ...interface{}) {}

// DebugKV 空操作
func (l *NoopLogger) DebugKV(msg string, keysAndValues ...interface{}) {}

// InfoKV 空操作
func (l *NoopLogger) InfoKV(msg string, keysAndValues ...interface{}) {}

// WarnKV 空操作
func (l *NoopLogger) WarnKV(msg string, keysAndValues ...interface{}) {}

// ErrorKV 空操作
func (l *NoopLogger) ErrorKV(msg string, keysAndValues ...interface{}) {}

// FatalKV 空操作
func (l *NoopLogger) FatalKV(msg string, keysAndValues ...interface{}) {}

// WithFields 返回当前实例
func (l *NoopLogger) WithFields(keysAndValues ...interface{}) Logger {
	return l
}

// Sync 空操作
func (l *NoopLogger) Sync() error {
	return nil
}
