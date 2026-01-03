package repository_impl

import (
	"context"

	"github.com/lyonnee/go-template/internal/infrastructure/cache"
	"github.com/lyonnee/go-template/internal/infrastructure/database"
)

type BaseRepository struct {
	db    database.DBContext
	cache cache.CacheContext
}

func (r *BaseRepository) SetContext(ctx context.Context) {
	r.cache, _ = cache.GetCacheContext(ctx)
	r.db, _ = database.GetDBContext(ctx)
}

func (r *BaseRepository) DB() database.DBContext {
	return r.db
}

func (r *BaseRepository) Cache() cache.CacheContext {
	return r.cache
}
