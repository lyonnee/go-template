package bootstrap

import (
	"github.com/lyonnee/go-template/interfaces/grpc"
	"github.com/lyonnee/go-template/interfaces/http"
)

func Initialize(env string) error {
	// config
	conf, err := initConfig(env)
	if err != nil {
		return err
	}

	// logger
	if _, err := initLogger(conf.Log); err != nil {
		return err
	}

	// database
	if _, err := initDatabase(conf.Persistence); err != nil {
		return err
	}

	// need init utils
	if err := initAuth(conf.Auth); err != nil {
		return err
	}

	if err := initUtils(conf.App.HostId); err != nil {
		return err
	}

	// business
	registerServices()

	// boot all servers
	go http.RunServer(conf.Http)
	go grpc.RunServer()

	return nil
}
