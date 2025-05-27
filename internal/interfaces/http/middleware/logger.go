package middleware

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
)

func Logger(logger log.Logger) app.HandlerFunc {
	return func(ctx context.Context, reqCtx *app.RequestContext) {
		start := time.Now() // 请求的时间

		reqCtx.Next(ctx) // 执行后续中间件

		cost := time.Since(start)

		logger.InfoKV("Request processed",
			"status", reqCtx.Response.StatusCode(), // 状态码
			"method", string(reqCtx.Request.Method()), // 请求的方法
			"path", string(reqCtx.Request.Path()), // 请求的路径
			"query", string(reqCtx.Request.QueryString()), // 请求的参数
			"ip", reqCtx.ClientIP(), // 请求的IP
			"user-agent", string(reqCtx.Request.Header.UserAgent()), // 请求头
			"errors", reqCtx.Errors.String(), // 错误信息
			"cost", cost, // 请求时间
			"trace_id", reqCtx.GetString("trace_id"), // 请求id
		)
	}
}
