package middleware

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lyonnee/go-template/bootstrap/di"
	"github.com/lyonnee/go-template/internal/infrastructure/auth"
	"github.com/lyonnee/go-template/internal/interfaces/http/dto"
)

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

		jwtManager := di.Get[*auth.JWTManager]()
		//解析token包含的信息
		claims, err := jwtManager.ValidateToken(parts[1])
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
