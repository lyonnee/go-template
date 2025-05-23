package config

import "time"

func Auth() AuthConfig {
	return conf.Auth
}

type AuthConfig struct {
	JWT JWTConfig `mapstructure:"jwt"`
}

type JWTConfig struct {
	SecretKey          string        `mapstructure:"secret_key"`           // 用于对 JWT 进行签名和验证的密钥
	AccessTokenExpiry  time.Duration `mapstructure:"access_token_expiry"`  // 访问令牌的有效时长（以秒为单位）
	RefreshTokenExpiry time.Duration `mapstructure:"refresh_token_expiry"` // 刷新令牌的有效时长（以秒为单位）
	Issuer             string        `mapstructure:"issuer"`               // 颁布单位
}
