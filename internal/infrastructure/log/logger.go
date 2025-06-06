package log

// Logger 定义通用日志接口，不依赖任何外部日志库
type Logger interface {
	// 基本日志方法
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})

	// 格式化日志方法
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})

	// 带键值对的日志方法
	DebugKV(msg string, keysAndValues ...interface{})
	InfoKV(msg string, keysAndValues ...interface{})
	WarnKV(msg string, keysAndValues ...interface{})
	ErrorKV(msg string, keysAndValues ...interface{})
	FatalKV(msg string, keysAndValues ...interface{})

	// 上下文方法
	WithFields(keysAndValues ...interface{}) Logger

	// 同步日志缓冲区
	Sync() error
}
