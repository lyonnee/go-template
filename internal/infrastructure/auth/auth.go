package auth

import "github.com/lyonnee/go-template/config"

var jwtManager *JWTManager

func Initialize(conf *config.AuthConfig) error {
	jwt := newJWTManger(&conf.JWT)
	jwtManager = jwt

	return nil
}

func JWT() *JWTManager {
	return jwtManager
}
