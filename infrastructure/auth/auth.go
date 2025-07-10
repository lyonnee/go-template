package auth

import (
	"github.com/lyonnee/go-template/infrastructure/config"
	"github.com/lyonnee/go-template/infrastructure/di"
)

var jwtManager *JWTManager

// func Initialize(conf config.AuthConfig) error {
// 	jwt := newJWTManager(&conf.JWT)
// 	jwtManager = jwt

// 	return nil
// }

func init() {
	conf := di.Get[config.Config]()
	jwtManager = newJWTManager(conf.Auth.JWT)
}

func JWTAuth() *JWTManager {
	return jwtManager
}
