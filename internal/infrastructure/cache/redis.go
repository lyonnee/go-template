// pkg/cache/redis.go
package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

// initRedis creates a new Redis cache instance
func initRedis(config RedisConfig) (*RedisCache, error) {
	var redisClient redis.UniversalClient

	if config.IsCluster() {
		redisClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    []string{config.Host + ":" + strconv.FormatInt(int64(config.Port), 10)},
			Username: config.Username,
			Password: config.Password,
		})
	} else {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     config.Host + ":" + strconv.FormatInt(int64(config.Port), 10),
			Username: config.Username,
			Password: config.Password,
			DB:       config.Database,
		})
	}

	// 测试连接
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisCache{
		redisCli:   redisClient,
		defaultTTL: config.TTL,
		prefix:     config.Prefix,
	}, nil
}

// RedisCache provides caching functionality using Redis
type RedisCache struct {
	prefix     string
	redisCli   redis.UniversalClient
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

// ZAdd adds members to a sorted set
func (r *RedisCache) ZAdd(ctx context.Context, key string, members ...ZMember) error {
	cacheKey := r.GetKey(key)

	zMembers := make([]*redis.Z, len(members))
	for i, m := range members {
		zMembers[i] = &redis.Z{
			Score:  m.Score,
			Member: m.Member,
		}
	}

	return r.redisCli.ZAdd(ctx, cacheKey, zMembers...).Err()
}

// ZRangeWithScores returns members in a sorted set by score (ascending)
func (r *RedisCache) ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]ZMember, error) {
	cacheKey := r.GetKey(key)

	result, err := r.redisCli.ZRangeWithScores(ctx, cacheKey, start, stop).Result()
	if err != nil {
		return nil, err
	}

	members := make([]ZMember, len(result))
	for i, z := range result {
		members[i] = ZMember{
			Score:  z.Score,
			Member: z.Member.(string),
		}
	}

	return members, nil
}

// ZRevRangeWithScores returns members in a sorted set by score (descending)
func (r *RedisCache) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]ZMember, error) {
	cacheKey := r.GetKey(key)

	result, err := r.redisCli.ZRevRangeWithScores(ctx, cacheKey, start, stop).Result()
	if err != nil {
		return nil, err
	}

	members := make([]ZMember, len(result))
	for i, z := range result {
		members[i] = ZMember{
			Score:  z.Score,
			Member: z.Member.(string),
		}
	}

	return members, nil
}

// ZCard returns the number of members in a sorted set
func (r *RedisCache) ZCard(ctx context.Context, key string) (int64, error) {
	cacheKey := r.GetKey(key)
	return r.redisCli.ZCard(ctx, cacheKey).Result()
}
