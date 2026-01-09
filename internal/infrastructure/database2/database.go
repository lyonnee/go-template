package database2

import (
	"context"
	"time"
)

type Config struct {
	Mysql    MysqlConfig    `yaml:"mysql"`
	Postgres PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	DSN string `yaml:"dsn"`

	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time"`
}

type MysqlConfig struct {
	DSN string `yaml:"dsn"`

	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time"`
}

type Executor interface {
	Executor() Executor
}

type Database interface {
	WithConnection(ctx context.Context, fn func(Executor) error) error
	WithTransaction(ctx context.Context, fn func(Executor) error) error
}

func init() {

}
