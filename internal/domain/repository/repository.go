package repository

import (
	"context"

	"github.com/lyonnee/go-template/internal/infrastructure/cache"
	"gorm.io/gorm"
)

type BaseRepository interface {
	SetContext(ctx context.Context)
	DB() *gorm.DB
	Cache() cache.CacheContext
}
