package cache

import (
	"context"
	"time"

	"github.com/lyonnee/go-template/pkg/di"
)

type Config struct {
	Redis RedisConfig `yaml:"redis"`
}

type RedisConfig struct {
	Host      string        `yaml:"host"`
	Port      int           `yaml:"port"`
	Username  string        `yaml:"username"`
	Password  string        `yaml:"password"`
	Database  int           `yaml:"database"`
	Framework string        `yaml:"framework"`
	Prefix    string        `yaml:"prefix"`
	TTL       time.Duration `yaml:"ttl"`
}

func (conf RedisConfig) IsCluster() bool {
	return conf.Framework == "cluster"
}

type Cache interface {
	GetKey(key string) string

	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string, dest any) error
	Delete(ctx context.Context, key string) error
	DeletePattern(ctx context.Context, pattern string) error
	Exists(ctx context.Context, key string) (bool, error)
	SetWithDefault(ctx context.Context, key string, value any) error
	Increment(ctx context.Context, key string) (int64, error)
	Decrement(ctx context.Context, key string) (int64, error)
	SetNX(ctx context.Context, key string, value any, ttl time.Duration) (bool, error)

	// Sorted Set 操作
	ZAdd(ctx context.Context, key string, members ...ZMember) error
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]ZMember, error)
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]ZMember, error)
	ZCard(ctx context.Context, key string) (int64, error)
}

// ZMember Sorted Set 成员
type ZMember struct {
	Score  float64
	Member string
}

var cache Cache

func init() {
	config := di.Get[*Config]()
	redisCache, err := initRedis(config.Redis)
	if err != nil {
		panic("Failed to initialize Redis client: " + err.Error())
	}

	cache = redisCache
	di.AddSingleton[Cache](func() (Cache, error) {
		return redisCache, nil
	})
}
