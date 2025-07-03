package bootstrap

import (
	stdLog "log"
	"os"

	"github.com/lyonnee/go-template/bootstrap/di"
	"github.com/lyonnee/go-template/infrastructure/config"
	"github.com/lyonnee/go-template/infrastructure/log"
	"go.uber.org/zap"
)

func initLogger(conf config.LogConfig) (*zap.Logger, error) {
	newLogger, err := log.Initialize(conf)
	if err != nil {
		stdLog.Printf("init logger failed, err:%s", err)
		os.Exit(1)
	}

	di.AddSingleton(func() (*zap.Logger, error) {
		return newLogger, nil
	})

	return newLogger, nil
}
