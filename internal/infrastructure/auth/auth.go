package auth

import (
	"github.com/lyonnee/go-template/internal/infrastructure/config"
	"github.com/lyonnee/go-template/pkg/di"
)

func init() {
	conf := di.Get[config.Config]()

	jwtGenerator := newJWTGenerator(conf.Auth.JWT)

	di.AddSingleton(func() (*JWTGenerator, error) {
		return jwtGenerator, nil
	})
}
