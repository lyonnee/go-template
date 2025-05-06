package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/lyonnee/go-template/pkg/modules/auth"
)

type Token struct {
	AccessToken  string `json:"access_token"`  // 用于身份验证，有效期较短（如 15 分钟）。
	RefreshToken string `json:"refresh_token"` // 用于刷新 Access Token，有效期较长（如 7 天），通常存储于安全位置（如 HttpOnly Cookie）
}

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		Claims:             auth.Claims{},
		KeyFunc:            auth.SecretKey(),
		TokenProcessorFunc: tokenProcessorFunc,
		SuccessHandler:     successHandler,
		ErrorHandler:       errorHandler,
		ContextKey:         "identity",
	})
}

func successHandler(*fiber.Ctx) error { return nil }

func tokenProcessorFunc(token string) (string, error) {
	return "", nil
}

func errorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
