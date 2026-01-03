// pkg/cache/redis.go
package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/lyonnee/go-template/internal/infrastructure/config"
	"github.com/lyonnee/go-template/pkg/di"
)

func init() {
	config := di.Get[config.Config]()
	redisCache, err := initRedis(config.Cache.Redis)
	if err != nil {
		panic("Failed to initialize Redis client: " + err.Error())
	}

	di.AddSingleton[CacheContext](func() (CacheContext, error) {
		return redisCache, nil
	})
}

// initRedis creates a new Redis cache instance
func initRedis(config config.RedisConfig) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + strconv.FormatInt(int64(config.Port), 10),
		Password: config.Password,
		DB:       config.Database,
	})

	// 测试连接
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisCache{
		redisCli:   client,
		defaultTTL: config.TTL,
	}, nil
}

// RedisCache provides caching functionality using Redis
type RedisCache struct {
	prefix     string
	redisCli   *redis.Client
	defaultTTL time.Duration
}

func (r *RedisCache) GetKey(key string) string {
	return fmt.Sprintf("%s:%s", r.prefix, key)
}

// Set stores a value in cache with specified TTL
func (r *RedisCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	cacheKey := r.GetKey(key)
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal cache value: %v", err)
	}

	if ttl == 0 {
		ttl = r.defaultTTL
	}

	return r.redisCli.Set(ctx, cacheKey, data, ttl).Err()
}

// Get retrieves a value from cache
func (r *RedisCache) Get(ctx context.Context, key string, dest any) error {
	cacheKey := r.GetKey(key)
	data, err := r.redisCli.Get(ctx, cacheKey).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), dest)
}

// Delete removes a key from cache
func (r *RedisCache) Delete(ctx context.Context, key string) error {
	cacheKey := r.GetKey(key)
	return r.redisCli.Del(ctx, cacheKey).Err()
}

// DeletePattern removes all keys matching a pattern
func (r *RedisCache) DeletePattern(ctx context.Context, pattern string) error {
	cachePattern := r.GetKey(pattern)
	keys, err := r.redisCli.Keys(ctx, cachePattern).Result()
	if err != nil {
		return err
	}

	if len(keys) > 0 {
		return r.redisCli.Del(ctx, keys...).Err()
	}

	return nil
}

// Exists checks if a key exists in cache
func (r *RedisCache) Exists(ctx context.Context, key string) (bool, error) {
	cacheKey := r.GetKey(key)
	count, err := r.redisCli.Exists(ctx, cacheKey).Result()
	return count > 0, err
}

// SetWithDefault stores a value in cache with default TTL
func (r *RedisCache) SetWithDefault(ctx context.Context, key string, value any) error {
	return r.Set(ctx, key, value, r.defaultTTL)
}

// Increment increments a numeric value in cache
func (r *RedisCache) Increment(ctx context.Context, key string) (int64, error) {
	cacheKey := r.GetKey(key)
	return r.redisCli.Incr(ctx, cacheKey).Result()
}

// Decrement decrements a numeric value in cache
func (r *RedisCache) Decrement(ctx context.Context, key string) (int64, error) {
	cacheKey := r.GetKey(key)
	return r.redisCli.Decr(ctx, cacheKey).Result()
}

// SetNX sets a key only if it doesn't exist
func (r *RedisCache) SetNX(ctx context.Context, key string, value any, ttl time.Duration) (bool, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return false, fmt.Errorf("failed to marshal cache value: %v", err)
	}

	if ttl == 0 {
		ttl = r.defaultTTL
	}

	cacheKey := r.GetKey(key)
	return r.redisCli.SetNX(ctx, cacheKey, data, ttl).Result()
}
