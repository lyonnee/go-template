package auth

import "github.com/lyonnee/go-template/infrastructure/config"

var jwtManager *JWTManager

func Initialize(conf config.AuthConfig) error {
	jwt := newJWTManger(&conf.JWT)
	jwtManager = jwt

	return nil
}

func JWTAuth() *JWTManager {
	return jwtManager
}
