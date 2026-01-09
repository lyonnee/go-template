package auth

import (
	"time"

	"github.com/lyonnee/go-template/pkg/di"
)

type Config struct {
	JWT JWTConfig `yaml:"jwt"`
}

type JWTConfig struct {
	SecretKey          string        `yaml:"secret_key"`           // 用于对 JWT 进行签名和验证的密钥
	AccessTokenExpiry  time.Duration `yaml:"access_token_expiry"`  // 访问令牌的有效时长（以秒为单位）
	RefreshTokenExpiry time.Duration `yaml:"refresh_token_expiry"` // 刷新令牌的有效时长（以秒为单位）
	Issuer             string        `yaml:"issuer"`               // 颁布单位
}

func init() {
	conf := di.Get[*Config]()

	jwtGenerator := newJWTGenerator(conf.JWT)

	di.AddSingleton(func() (*JWTGenerator, error) {
		return jwtGenerator, nil
	})
}
