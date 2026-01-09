package repository_impl

import (
	"context"

	"github.com/lyonnee/go-template/internal/infrastructure/cache"
	"github.com/lyonnee/go-template/internal/infrastructure/database"
	"gorm.io/gorm"
)

type BaseRepository struct {
	ctx context.Context
}

func (r *BaseRepository) SetContext(ctx context.Context) {
	r.SetContext(ctx)
}

func (r *BaseRepository) DB() (*gorm.DB, error) {
	return database.GetDBContext(r.ctx)
}

func (r *BaseRepository) Cache() (cache.Cache, error) {
	return cache.GetCacheContext(r.ctx)
}
