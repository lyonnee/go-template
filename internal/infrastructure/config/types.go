package config

import "time"

// ================== AppConfig ==================
type AppConfig struct {
	Env         string `mapstructure:"env"`         // 环境变量
	Name        string `mapstructure:"name"`        // 应用名称
	Version     string `mapstructure:"version"`     // 应用版本
	Description string `mapstructure:"description"` // 应用描述
	HostId      int64  `mapstructure:"host_id"`     // 主机id
}

// ================== AuthConfig ==================
type AuthConfig struct {
	JWT JWTConfig `mapstructure:"jwt"`
}

type JWTConfig struct {
	SecretKey          string        `mapstructure:"secret_key"`           // 用于对 JWT 进行签名和验证的密钥
	AccessTokenExpiry  time.Duration `mapstructure:"access_token_expiry"`  // 访问令牌的有效时长（以秒为单位）
	RefreshTokenExpiry time.Duration `mapstructure:"refresh_token_expiry"` // 刷新令牌的有效时长（以秒为单位）
	Issuer             string        `mapstructure:"issuer"`               // 颁布单位
}

// ================== DatabaseConfig ==================
type DatabaseConfig struct {
	Mysql    MysqlConfig    `mapstructure:"mysql"`
	Postgres PostgresConfig `mapstructure:"postgres"`
}

type PostgresConfig struct {
	DSN string `mapstructure:"dsn"`

	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
}

type MysqlConfig struct {
	DSN string `mapstructure:"dsn"`

	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
}

// ================== CacheConfig ==================
type CacheConfig struct {
	Redis RedisConfig `mapstructure:"redis"`
}

type RedisConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Database  int    `mapstructure:"database"`
	Framework string `mapstructure:"framework"`
	Prefix    string `mapstructure:"prefix"`
}

func (conf RedisConfig) IsCluster() bool {
	return conf.Framework == "cluster"
}

// ================== HttpConfig ==================
// HttpConfig 包含 HTTP 服务的配置

type HttpConfig struct {
	Port string `mapstructure:"port"`
}

// ==================  LogConfig ==================
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
