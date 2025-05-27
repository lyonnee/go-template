package config

func Log() LogConfig {
	return conf.Log
}

type LogConfig struct {
	Format string `mapstructure:"format"`
	Caller string `mapstructure:"caller"`

	Level  string `mapstructure:"level"`
	ToFile bool   `mapstructure:"to_file"`
	// 日志文件配置
	LogFileConfig LogFileConfig `mapstructure:"log_file_config"`
}

type LogFileConfig struct {
	Filename      string `mapstructure:"filename"`
	LogLevel      string `mapstructure:"log_level"`
	MaxSize       int    `mapstructure:"max_size"`
	MaxAge        int    `mapstructure:"max_age"`
	MaxBackups    int    `mapstructure:"max_backups"`
	IsCompression bool   `mapstructure:"is_compression"`
}
