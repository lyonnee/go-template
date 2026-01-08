package repository

import (
	"context"

	"github.com/lyonnee/go-template/internal/infrastructure/cache"
	"gorm.io/gorm"
)

type BaseRepository interface {
	SetContext(ctx context.Context)
	DB() (*gorm.DB, error)
	Cache() (cache.CacheContext, error)
}
