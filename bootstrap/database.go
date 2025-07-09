package bootstrap

import (
	"github.com/lyonnee/go-template/bootstrap/di"
	"github.com/lyonnee/go-template/infrastructure/config"
	"github.com/lyonnee/go-template/infrastructure/database"
	"go.uber.org/zap"
)

func initDatabase(conf config.PersistenceConfig) (database.Database, error) {
	logger := di.Get[*zap.Logger]()
	logger = logger.WithOptions(zap.WithCaller(false))

	db, err := database.Initialize(conf, logger)
	if err != nil {
		return nil, err
	}

	di.AddSingleton(func() (database.Database, error) {
		return db, nil
	})

	return db, nil
}
