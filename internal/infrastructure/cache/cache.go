package cache

import (
	"time"

	"github.com/lyonnee/go-template/pkg/di"
)

type Config struct {
	Redis RedisConfig `mapstructure:"redis"`
}

type RedisConfig struct {
	Host      string        `mapstructure:"host"`
	Port      int           `mapstructure:"port"`
	Username  string        `mapstructure:"username"`
	Password  string        `mapstructure:"password"`
	Database  int           `mapstructure:"database"`
	Framework string        `mapstructure:"framework"`
	Prefix    string        `mapstructure:"prefix"`
	TTL       time.Duration `mapstructure:"ttl"`
}

func (conf RedisConfig) IsCluster() bool {
	return conf.Framework == "cluster"
}

func init() {
	config := di.Get[*Config]()
	redisCache, err := initRedis(config.Redis)
	if err != nil {
		panic("Failed to initialize Redis client: " + err.Error())
	}

	di.AddSingleton[CacheContext](func() (CacheContext, error) {
		return redisCache, nil
	})
}
