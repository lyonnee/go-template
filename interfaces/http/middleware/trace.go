package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lyonnee/go-template/pkg/idgen"
)

func AddTrace() app.HandlerFunc {
	return func(ctx context.Context, reqCtx *app.RequestContext) {
		// 从请求头获取或生成 Trace ID
		traceID := string(reqCtx.GetHeader("X-Trace-ID"))
		if traceID == "" {
			traceID = GenerateTraceID() // 使用前述方法生成
		}

		// 存入上下文
		reqCtx.Set("trace_id", traceID)

		// 设置响应头（可选）
		reqCtx.Header("X-Trace-ID", traceID)
		reqCtx.Next(ctx)
	}
}

func GenerateTraceID() string {
	return idgen.GenerateStringId()
}
