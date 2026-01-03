package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lyonnee/go-template/pkg/idgen"
)

func AddTrace() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取或生成 Trace ID
		traceID := c.GetHeader("X-Trace-ID")
		if traceID == "" {
			traceID = GenerateTraceID() // 使用前述方法生成
		}

		// 存入上下文
		c.Set("trace_id", traceID)

		// 设置响应头（可选）
		c.Header("X-Trace-ID", traceID)
		c.Next()
	}
}

func GenerateTraceID() string {
	return idgen.GenerateStringId()
}
