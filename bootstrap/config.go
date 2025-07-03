package bootstrap

import (
	"log"
	"os"

	"github.com/lyonnee/go-template/infrastructure/config"
)

func initConfig(env string) (*config.Config, error) {
	conf, err := config.Load(env)
	if err != nil {
		log.Printf("load config failed, err:%s", err)
		os.Exit(1)

	}

	return conf, nil
}
