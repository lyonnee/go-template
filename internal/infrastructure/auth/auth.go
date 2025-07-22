package auth

import (
	"github.com/lyonnee/go-template/internal/infrastructure/config"
	"github.com/lyonnee/go-template/internal/infrastructure/di"
)

var jwtManager *JWTManager

func init() {
	conf := di.Get[config.Config]()

	jwtManager := newJWTManager(conf.Auth.JWT)

	di.AddSingleton(func() (*JWTManager, error) {
		return jwtManager, nil
	})
}
