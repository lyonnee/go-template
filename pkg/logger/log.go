package logger

import (
	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/lyonnee/go-template/pkg/container"
)

func Initialize() error {
	config := container.GetService[*config.Config]().Log

	logger, err := newLogger(config)
	if err != nil {
		return err
	}

	container.AddSingletonService[log.Logger](func() (log.Logger, error) {
		return logger, nil
	})

	return nil
}

func newLogger(config config.LogConfig) (log.Logger, error) {

	zapLogger, err := NewZapLogger(config)
	if err != nil {
		return nil, err
	}
	logger := NewZapSugarLogger(zapLogger)

	return logger, nil
}
