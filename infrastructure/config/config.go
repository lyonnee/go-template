package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/lyonnee/go-template/infrastructure/di"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Http     HttpConfig     `mapstructure:"http"`
	Log      LogConfig      `mapstructure:"log"`
	Auth     AuthConfig     `mapstructure:"auth"`
	Database DatabaseConfig `mapstructure:"database"`
	Cache    CacheConfig    `mapstructure:"cache"`
}

var conf = new(Config)

func init() {
	var (
		env = flag.String("env", "dev", "Environment (dev, test, prod)")
	)
	flag.Parse()

	Load(*env)
}

func Load(env string) (*Config, error) {
	if env == "" {
		env = "prod"
	}

	// 使用viper作为配置加载中间件
	workDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir)

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			return nil, fmt.Errorf("config file not found for environment %s: %w", env, err)
		default:
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	if err := viper.Unmarshal(conf); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	di.AddSingleton[Config](func() (Config, error) {
		return *conf, nil
	})

	return conf, nil
}
