package config

func Log() LogConfig {
	return conf.Log
}

type LogConfig struct {
	// console
	EnableToConsole bool   `mapstructure:"enable_to_console"`
	ToConsoleLevel  string `mapstructure:"to_console_level"`

	// file
	ToFileLevel string `mapstructure:"to_file_level"`
	Filename    string `mapstructure:"filename"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxAge      int    `mapstructure:"max_age"`
	MaxBackups  int    `mapstructure:"max_backups"`
}
