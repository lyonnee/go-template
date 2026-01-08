package database

import (
	"context"
	"time"

	"github.com/lyonnee/go-template/pkg/di"
	"github.com/lyonnee/go-template/pkg/log"
	"gorm.io/gorm"
)

type Config struct {
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

type Database struct {
	db *gorm.DB
}

func (dbc *Database) WithContext(ctx context.Context, fn func(context.Context) error) error {
	return fn(SetDBContext(ctx, dbc.db.WithContext(ctx)))
}

func (dbc *Database) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	return dbc.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(SetDBContext(ctx, tx))
	})
}

func (dbc *Database) Close() error {
	if dbc.db != nil {
		sqlDB, err := dbc.db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

func (dbc *Database) CloseWithContext(ctx context.Context) error {
	if dbc.db == nil {
		return nil
	}
	done := make(chan error, 1)
	go func() {
		done <- dbc.Close()
	}()
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

var db *Database

func init() {
	config := di.Get[*Config]()
	logger := di.Get[*log.Logger]()

	gormDB, err := newGormDB(config.Postgres, logger)
	if err != nil {
		panic("Failed to initialize PostgreSQL database: " + err.Error())
	}

	db = gormDB
	di.AddSingleton[DBContext](func() (DBContext, error) {
		return db, nil
	})
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// CloseWithContext closes the global DB instance with context awareness.
func CloseWithContext(ctx context.Context) error {
	if db != nil {
		return db.CloseWithContext(ctx)
	}
	return nil
}
