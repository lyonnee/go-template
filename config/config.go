package config

import (
	"fmt"
	"os"

	"github.com/lyonnee/go-template/pkg/container"
	"github.com/spf13/viper"
)

type Config struct {
	Http        HttpConfig        `mapstructure:"http"`
	Log         LogConfig         `mapstructure:"log"`
	Auth        AuthConfig        `mapstructure:"auth"`
	Persistence PersistenceConfig `mapstructure:"persistence"`
	Cache       CacheConfig       `mapstructure:"cache"`
}

var conf = new(Config)

func Load(env string) error {
	if env == "" {
		env = "prod"
	}

	// 使用viper作为配置加载中间件
	workDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir)

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			return fmt.Errorf("config file not found for environment %s: %w", env, err)
		default:
			return fmt.Errorf("error reading config file: %w", err)
		}
	}

	if err := viper.Unmarshal(conf); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	container.AddSingletonService[*Config](func() (*Config, error) {
		return conf, nil
	})

	return nil
}
