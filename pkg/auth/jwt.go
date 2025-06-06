package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/pkg/container"
)

// Claims 自定义Claims
type Claims struct {
	UserId        int64  `json:"user_id"`
	AlternativeID string `json:"alternative_id"` // 备选id， username/email/phone number/wallet address
	jwt.RegisteredClaims
}

func SecretKey() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		config := container.GetService[*config.Config]().Auth

		return []byte(config.JWT.SecretKey), nil
	}
}

// GenerateAccessToken 构建访问token
// @dev AccessToken 用于身份验证，有效期较短（如 15 分钟）
func GenerateAccessToken(userID int64, alternativeID string) (string, error) {
	config := container.GetService[*config.Config]().Auth

	return genToken(userID, alternativeID, config.JWT.AccessTokenExpiry)
}

// GenerateRefreshToken 生成刷新令牌
// @dev RefreshToken 用于刷新 Access Token，有效期较长（如 7 天），通常存储于安全位置（如 HttpOnly Cookie）
func GenerateRefreshToken(userID int64, alternativeID string) (string, error) {
	config := container.GetService[*config.Config]().Auth

	return genToken(userID, alternativeID, config.JWT.RefreshTokenExpiry)
}

// RefreshToken 刷新JWT（使用刷新令牌生成新的访问令牌）
func RefreshToken(refreshToken string) (string, error) {
	// 验证刷新令牌
	claims, err := ValidateToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("refresh token is invalid or expired: %w", err)
	}

	// 使用刷新令牌中的用户名生成新的访问令牌
	return GenerateAccessToken(claims.UserId, claims.AlternativeID)
}

// ValidateToken 验证JWT令牌
func ValidateToken(tokenString string) (*Claims, error) {
	config := container.GetService[*config.Config]().Auth

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(config.JWT.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// 验证Issuer
		if claims.Issuer != config.JWT.Issuer {
			return nil, fmt.Errorf("invalid issuer")
		}
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func genToken(userID int64, alternativeID string, expiry time.Duration) (string, error) {
	config := container.GetService[*config.Config]().Auth

	claims := Claims{
		AlternativeID: alternativeID,
		UserId:        userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(expiry),
			),
			Issuer: config.JWT.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT.SecretKey))
}
