package config

import "time"

// ======================================== Base Config ============================================= //

type Config struct {
	Http        HttpConfig        `yaml:"http"`
	Log         LogConfig         `yaml:"log"`
	Persistence PersistenceConfig `yaml:"persistence"`
	Auth        AuthConfig        `yaml:"auth"`
}

// =================================================================================================== //

type HttpConfig struct {
	Port string
}

type AuthConfig struct {
	JWT JWTConfig `yaml:"jwt"`
}

type JWTConfig struct {
	SecretKey          string        `yaml:"secret_key"`           // 用于对 JWT 进行签名和验证的密钥
	AccessTokenExpiry  time.Duration `yaml:"access_token_expiry"`  // 访问令牌的有效时长（以秒为单位）
	RefreshTokenExpiry time.Duration `yaml:"refresh_token_expiry"` // 刷新令牌的有效时长（以秒为单位）
	Issuer             string        `yaml:"issuer"`               // 颁布单位
}

type LogConfig struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

type PersistenceConfig struct {
	Mysql    MysqlConfig    `yaml:"mysql"`
	Postgres PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	DSN string
}

type MysqlConfig struct {
	DSN string
}
