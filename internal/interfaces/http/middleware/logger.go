package middleware

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lyonnee/go-template/pkg/log"
	"go.uber.org/zap"
)

func Logger() app.HandlerFunc {
	return func(ctx context.Context, reqCtx *app.RequestContext) {
		start := time.Now() // 请求的时间

		reqCtx.Next(ctx) // 执行后续中间件

		cost := time.Since(start)

		log.Info("",
			zap.Int("status", reqCtx.Response.StatusCode()),                     // 状态码
			zap.String("method", string(reqCtx.Request.Method())),               // 请求的方法
			zap.String("path", string(reqCtx.Request.Path())),                   // 请求的路径
			zap.String("query", string(reqCtx.Request.QueryString())),           // 请求的参数
			zap.String("ip", reqCtx.ClientIP()),                                 // 请求的IP
			zap.String("user-agent", string(reqCtx.Request.Header.UserAgent())), // 请求头
			zap.String("errors", reqCtx.Errors.String()),                        // 错误信息
			zap.Duration("cost", cost),                                          // 请求时间
			zap.String("trace_id", reqCtx.GetString("trace_id")),                // 请求id
		)
	}
}
