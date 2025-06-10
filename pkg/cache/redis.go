// pkg/cache/redis.go
package cache

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/lyonnee/go-template/config"
)

// InitRedis 初始化Redis客户端
func InitRedis(config config.RedisConfig) (*redis.Client, error) {
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
