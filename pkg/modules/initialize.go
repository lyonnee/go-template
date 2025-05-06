package modules

import (
	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/pkg/modules/logger"
	"github.com/lyonnee/go-template/pkg/modules/persistence"
)

func Initialize() error {
	logger.Initialize(config.Log())

	persistence.Initialize(config.Persistence())

	return nil
}
