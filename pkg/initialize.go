package pkg

import (
	"github.com/lyonnee/go-template/pkg/cache"
	"github.com/lyonnee/go-template/pkg/logger"
	"github.com/lyonnee/go-template/pkg/mq"
	"github.com/lyonnee/go-template/pkg/persistence"
)

func Initialize() error {
	if err := logger.Initialize(); err != nil {
		return err
	}

	if err := mq.Initialize(); err != nil {
		return err
	}

	if err := cache.Initialize(); err != nil {
		return err
	}

	if err := persistence.Initialize(); err != nil {
		return err
	}

	return nil
}
