package log

import (
	"github.com/lyonnee/go-template/pkg/di"
	"go.uber.org/zap"
)

type Config struct {
	// 控制台配置
	ConsoleWriterConfig LogConsoleWriterConfig `mapstructure:"console_writer_config"`
	// 日志文件配置
	FileWriterConfig LogFileWriterConfig `mapstructure:"file_writer_config"`
}

type LogConsoleWriterConfig struct {
	Enable bool `mapstructure:"enable"`

	Format string `mapstructure:"format"`
	Level  string `mapstructure:"level"`
	Caller string `mapstructure:"caller"`
}

type LogFileWriterConfig struct {
	Enable bool `mapstructure:"enable"`

	Format   string `mapstructure:"format"`
	Filename string `mapstructure:"filename"`
	Level    string `mapstructure:"level"`
	Caller   string `mapstructure:"caller"`

	MaxSize       int  `mapstructure:"max_size"`
	MaxAge        int  `mapstructure:"max_age"`
	MaxBackups    int  `mapstructure:"max_backups"`
	IsCompression bool `mapstructure:"is_compression"`
}

type Logger = zap.Logger

var (
	logger *Logger
)

func init() {
	config := di.Get[*Config]()
	newLogger, err := newZapLogger(config)
	if err != nil {
		panic(err)
	}

	logger = newLogger
	di.AddSingleton[Logger](func() (Logger, error) {
		return *logger, nil
	})
}

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
