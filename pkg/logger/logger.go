package logger

import (
	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
)

func NewLogger(config config.LogConfig) (log.Logger, error) {
	return newLogger(config)
}

func newLogger(config config.LogConfig) (log.Logger, error) {

	zapLogger, err := NewZapLogger(config)
	if err != nil {
		return nil, err
	}
	logger := NewZapSugarLogger(zapLogger)

	return logger, nil
}
