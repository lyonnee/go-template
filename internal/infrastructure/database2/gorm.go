package database2

import (
	"context"

	"github.com/lyonnee/go-template/pkg/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormExecutor struct {
	*gorm.DB
	isTransaction bool
}

type GormDB struct {
	*gorm.DB
}

func (executor *GormExecutor) Executor() Executor {
	if executor.isTransaction {
		return executor
	}

	newExecutor := executor.Session(&gorm.Session{NewDB: true})
	return &GormExecutor{DB: newExecutor}
}

func (db *GormDB) WithConnection(ctx context.Context, fn func(Executor) error) error {
	GormExecutor := &GormExecutor{DB: db.WithContext(ctx), isTransaction: true}
	return fn(GormExecutor)
}

func (db *GormDB) WithTransaction(ctx context.Context, fn func(Executor) error) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		GormExecutor := &GormExecutor{DB: tx, isTransaction: true}
		return fn(GormExecutor)
	})
}

func newGormDB(cfg PostgresConfig, logger log.Logger) (*GormDB, error) {
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

	return &GormDB{db}, nil
}
