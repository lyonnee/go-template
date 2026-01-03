package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lyonnee/go-template/pkg/log"
	"go.uber.org/zap"
)

func Logger(logger *log.Logger) gin.HandlerFunc {
	logger = logger.WithOptions(zap.WithCaller(false)) // Skip the logger call in the stack trace

	return func(c *gin.Context) {
		start := time.Now() // 请求的时间

		c.Next() // 执行后续中间件

		cost := time.Since(start)

		logger.Info("http request",
			zap.Int("status", c.Writer.Status()),           // 状态码
			zap.String("method", c.Request.Method),         // 请求的方法
			zap.String("path", c.Request.URL.Path),         // 请求的路径
			zap.String("query", c.Request.URL.RawQuery),    // 请求的参数
			zap.String("ip", c.ClientIP()),                 // 请求的IP
			zap.String("user-agent", c.Request.UserAgent()), // 请求头
			zap.String("errors", c.Errors.String()),        // 错误信息
			zap.String("cost", cost.String()),              // 请求时间
			zap.String("trace_id", c.GetString("trace_id")), // 请求id
		)
	}
}
