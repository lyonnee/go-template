package middleware

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lyonnee/go-template/internal/interfaces/http/dto"
	"github.com/lyonnee/go-template/pkg/auth"
)

type Token struct {
	AccessToken  string `json:"access_token"`  // 用于身份验证，有效期较短（如 15 分钟）。
	RefreshToken string `json:"refresh_token"` // 用于刷新 Access Token，有效期较长（如 7 天），通常存储于安全位置（如 HttpOnly Cookie）
}

// JWTAuth 中间件，检查token
func JWTAuth() app.HandlerFunc {
	return func(ctx context.Context, reqCtx *app.RequestContext) {
		authHeader := reqCtx.Request.Header.Get("Authorization")
		if authHeader == "" {
			dto.Fail(reqCtx, dto.CODE_NOT_TOKEN, "Access Denied. Token not included in the request.")
			reqCtx.Abort() //结束后续操作
			return
		}

		//按空格拆分
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			dto.Fail(reqCtx, dto.CODE_TOKEN_FORMAT_INCORRECT, "The format of the auth in the request header is incorrect.")
			reqCtx.Abort()
			return
		}

		//解析token包含的信息
		claims, err := auth.ValidateToken(parts[1])
		if err != nil {
			dto.Fail(reqCtx, dto.CODE_TOKEN_INVALID, "Invalid JSON Web Token")
			reqCtx.Abort()
			return
		}

		// 将当前请求的claims信息保存到请求的上下文c上
		reqCtx.Set("claims", claims)
		reqCtx.Next(ctx) // 后续的处理函数可以用过ctx.Get("claims")来获取当前请求的用户信息
	}
}
