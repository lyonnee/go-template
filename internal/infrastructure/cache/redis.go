// pkg/cache/redis.go
package cache

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/lyonnee/go-template/internal/infrastructure/config"
	"github.com/lyonnee/go-template/internal/infrastructure/di"
)

func init() {
	config := di.Get[config.Config]()
	redisClient, err := initRedis(config.Cache.Redis)
	if err != nil {
		panic("Failed to initialize Redis client: " + err.Error())
	}

	di.AddSingleton[*redis.Client](func() (*redis.Client, error) {
		return redisClient, nil
	})
}

// initRedis 初始化Redis客户端
func initRedis(config config.RedisConfig) (*redis.Client, error) {
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

	return client, nil
}
