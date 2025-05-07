package config

import "time"

// ======================================== Base Config ============================================= //

type Config struct {
	Http        HttpConfig        `mapstructure:"http"`
	Log         LogConfig         `mapstructure:"log"`
	Persistence PersistenceConfig `mapstructure:"persistence"`
	Auth        AuthConfig        `mapstructure:"auth"`
}

// =================================================================================================== //

type HttpConfig struct {
	Port string `mapstructure:"port"`
}

type AuthConfig struct {
	JWT JWTConfig `mapstructure:"jwt"`
}

type JWTConfig struct {
	SecretKey          string        `mapstructure:"secret_key"`           // 用于对 JWT 进行签名和验证的密钥
	AccessTokenExpiry  time.Duration `mapstructure:"access_token_expiry"`  // 访问令牌的有效时长（以秒为单位）
	RefreshTokenExpiry time.Duration `mapstructure:"refresh_token_expiry"` // 刷新令牌的有效时长（以秒为单位）
	Issuer             string        `mapstructure:"issuer"`               // 颁布单位
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

type PersistenceConfig struct {
	Mysql    MysqlConfig    `mapstructure:"mysql"`
	Postgres PostgresConfig `mapstructure:"postgres"`
}

type PostgresConfig struct {
	DSN string `mapstructure:"dsn"`
}

type MysqlConfig struct {
	DSN string `mapstructure:"dsn"`
}
