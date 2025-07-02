package bootstrap

import (
	"github.com/lyonnee/go-template/internal/interfaces/grpc"
	"github.com/lyonnee/go-template/internal/interfaces/http"
)

func Initialize(env string) error {
	// config
	conf, err := initConfig(env)
	if err != nil {
		return err
	}

	// logger
	logger, err := initLogger(conf.Log)
	if err != nil {
		return err
	}

	// data access
	if _, err := initPersistence(conf.Persistence, logger); err != nil {
		return err
	}

	// need init utils
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
