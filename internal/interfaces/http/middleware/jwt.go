package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lyonnee/go-template/internal/infrastructure/auth"
	"github.com/lyonnee/go-template/internal/interfaces/http/dto"
	"github.com/lyonnee/go-template/pkg/di"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			dto.Fail(c, dto.CODE_NOT_TOKEN, "Access Denied. Token not included in the request.")
			c.Abort() //结束后续操作
			return
		}

		//按空格拆分
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			dto.Fail(c, dto.CODE_TOKEN_FORMAT_INCORRECT, "The format of the auth in the request header is incorrect.")
			c.Abort()
			return
		}

		jwtGenerator := di.Get[*auth.JWTGenerator]()
		//解析token包含的信息
		claims, err := jwtGenerator.ValidateToken(parts[1])
		if err != nil {
			dto.Fail(c, dto.CODE_TOKEN_INVALID, "Invalid JSON Web Token")
			c.Abort()
			return
		}

		// 将当前请求的claims信息保存到请求的上下文c上
		c.Set("claims", claims)
		c.Next() // 后续的处理函数可以用过c.Get("claims")来获取当前请求的用户信息
	}
}
