package config

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/lyonnee/go-template/internal/infrastructure/auth"
	"github.com/lyonnee/go-template/internal/infrastructure/cache"
	"github.com/lyonnee/go-template/internal/infrastructure/database"
	"github.com/lyonnee/go-template/pkg/di"
	"github.com/lyonnee/go-template/pkg/log"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig       `yaml:"app"`
	Http     HttpConfig      `yaml:"http"`
	Grpc     GRPCConfig      `yaml:"grpc"`
	Log      log.Config      `yaml:"log"`
	Auth     auth.Config     `yaml:"auth"`
	Database database.Config `yaml:"database"`
	Cache    cache.Config    `yaml:"cache"`
}

var conf = new(Config)

func init() {
	var (
		env = flag.String("env", "dev", "Environment (dev, test, prod)")
	)
	flag.Parse()

	newConf, err := Load(*env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	conf = newConf

	di.AddSingleton[Config](func() (Config, error) {
		return *conf, nil
	})

	di.AddSingleton[*log.Config](func() (*log.Config, error) {
		return &conf.Log, nil
	})

	di.AddSingleton[*auth.Config](func() (*auth.Config, error) {
		return &conf.Auth, nil
	})

	di.AddSingleton[*database.Config](func() (*database.Config, error) {
		return &conf.Database, nil
	})

	di.AddSingleton[*cache.Config](func() (*cache.Config, error) {
		return &conf.Cache, nil
	})
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
	viper.AddConfigPath(path.Join(workDir, "configs"))

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			return nil, fmt.Errorf("config file not found for environment %s: %w", env, err)
		default:
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var newConf Config
	if err := viper.Unmarshal(&newConf); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &newConf, nil
}
