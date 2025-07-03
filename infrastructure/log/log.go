package log

import (
	"github.com/lyonnee/go-template/infrastructure/config"
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
)

func Initialize(logConfig config.LogConfig,
) (*zap.Logger, error) {
	logger, err := newZapLogger(logConfig)
	if err != nil {
		return nil, err
	}
	Logger = logger

	return Logger, nil
}
