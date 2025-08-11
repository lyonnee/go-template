package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lyonnee/go-template/internal/infrastructure/config"
)

// Claims 自定义Claims
type Claims struct {
	UserId        uint64 `json:"user_id"`
	AlternativeID string `json:"alternative_id"` // 备选id， username/email/phone number/wallet address
	jwt.RegisteredClaims
}

type JWTGenerator struct {
	secretKey          []byte
	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
	issuer             string
}

func newJWTGenerator(conf config.JWTConfig) *JWTGenerator {
	jwtManager := &JWTGenerator{
		secretKey:          []byte(conf.SecretKey),
		accessTokenExpiry:  conf.AccessTokenExpiry,
		refreshTokenExpiry: conf.RefreshTokenExpiry,
		issuer:             conf.Issuer,
	}

	return jwtManager
}

func (m *JWTGenerator) SecretKey() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return m.secretKey, nil
	}
}

// GenerateAccessToken 构建访问token
// @dev AccessToken 用于身份验证，有效期较短（如 15 分钟）
func (m *JWTGenerator) GenerateAccessToken(userID uint64, alternativeID string) (string, error) {
	return m.genToken((userID), alternativeID, m.accessTokenExpiry)
}

// GenerateRefreshToken 生成刷新令牌
// @dev RefreshToken 用于刷新 Access Token，有效期较长（如 7 天），通常存储于安全位置（如 HttpOnly Cookie）
func (m *JWTGenerator) GenerateRefreshToken(userID uint64, alternativeID string) (string, error) {
	return m.genToken((userID), alternativeID, m.refreshTokenExpiry)
}

// RefreshToken 刷新JWT（使用刷新令牌生成新的访问令牌）
func (m *JWTGenerator) RefreshToken(refreshToken string) (string, error) {
	// 验证刷新令牌
	claims, err := m.ValidateToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("refresh token is invalid or expired: %w", err)
	}

	// 使用刷新令牌中的用户名生成新的访问令牌
	return m.GenerateAccessToken(claims.UserId, claims.AlternativeID)
}

// ValidateToken 验证JWT令牌
func (m *JWTGenerator) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return m.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// 验证Issuer
		if claims.Issuer != m.issuer {
			return nil, fmt.Errorf("invalid issuer")
		}
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func (m *JWTGenerator) genToken(userID uint64, alternativeID string, expiry time.Duration) (string, error) {
	claims := Claims{
		AlternativeID: alternativeID,
		UserId:        userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(expiry),
			),
			Issuer: m.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.secretKey)
}
