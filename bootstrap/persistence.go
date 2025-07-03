package bootstrap

import (
	"github.com/lyonnee/go-template/bootstrap/di"
	"github.com/lyonnee/go-template/infrastructure/config"
	"github.com/lyonnee/go-template/infrastructure/persistence"
	"go.uber.org/zap"
)

func initPersistence(conf config.PersistenceConfig, logger *zap.Logger) (persistence.DBContext, error) {
	db, err := persistence.Initialize(conf, logger)
	if err != nil {
		return nil, err
	}

	di.AddSingleton(func() (persistence.DBContext, error) {
		return db, nil
	})

	return db, nil
}
