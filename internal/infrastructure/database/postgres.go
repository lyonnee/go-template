package database

import (
	"github.com/lyonnee/go-template/internal/infrastructure/config"
	"github.com/lyonnee/go-template/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newPostgresDB(cfg config.PostgresConfig, logger *log.Logger) (*Database, error) {
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
