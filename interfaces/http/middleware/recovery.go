package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go.uber.org/zap"
)

func Recovery(logger *zap.Logger) func(ctx context.Context, reqCtx *app.RequestContext, err interface{}, stack []byte) {
	return func(ctx context.Context, reqCtx *app.RequestContext, err interface{}, stack []byte) {
		logger.Fatal("[Recovery from panic]", zap.Any("err", err), zap.ByteString("stack", stack))
		reqCtx.AbortWithStatus(consts.StatusInternalServerError)
	}
}
