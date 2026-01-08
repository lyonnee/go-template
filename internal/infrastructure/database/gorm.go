package database

import (
	"context"

	"github.com/lyonnee/go-template/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDB struct {
	db *gorm.DB
}

func (dbc *GormDB) WithContext(ctx context.Context, fn func(context.Context) error) error {
	return fn(SetDBContext(ctx, dbc.db.WithContext(ctx)))
}
func (dbc *GormDB) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	return dbc.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(SetDBContext(ctx, tx))
	})
}

func newGormDB(cfg PostgresConfig, logger *log.Logger) (*Database, error) {
	gormConfig := &gorm.Config{
		Logger: NewGormLogger(logger),
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN), gormConfig)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	return &Database{db: db}, nil
}
