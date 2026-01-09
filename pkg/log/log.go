package log

import (
	"github.com/lyonnee/go-template/pkg/di"
)

type Config struct {
	// 控制台配置
	ConsoleWriterConfig LogConsoleWriterConfig `yaml:"console_writer_config"`
	// 日志文件配置
	FileWriterConfig LogFileWriterConfig `yaml:"file_writer_config"`
}

type LogConsoleWriterConfig struct {
	Enable bool `yaml:"enable"`

	Format string `yaml:"format"`
	Level  string `yaml:"level"`
	Caller string `yaml:"caller"`
}

type LogFileWriterConfig struct {
	Enable bool `yaml:"enable"`

	Format   string `yaml:"format"`
	Filename string `yaml:"filename"`
	Level    string `yaml:"level"`
	Caller   string `yaml:"caller"`

	MaxSize       int  `yaml:"max_size"`
	MaxAge        int  `yaml:"max_age"`
	MaxBackups    int  `yaml:"max_backups"`
	IsCompression bool `yaml:"is_compression"`
}

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Close() error
}

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
