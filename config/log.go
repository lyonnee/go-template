package config

type LogConfig struct {
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
