package repository

import (
	"context"

	"github.com/lyonnee/go-template/internal/infrastructure/cache"
	"github.com/lyonnee/go-template/internal/infrastructure/database"
)

type BaseRepository interface {
	SetContext(ctx context.Context)
	DB() database.DBContext
	Cache() cache.CacheContext
}
