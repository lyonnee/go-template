package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
)

func Recovery(logger log.Logger) func(ctx context.Context, reqCtx *app.RequestContext, err interface{}, stack []byte) {
	return func(ctx context.Context, reqCtx *app.RequestContext, err interface{}, stack []byte) {
		logger.FatalKV("[Recovery from panic]", "err", err, "stack", string(stack))
		reqCtx.AbortWithStatus(consts.StatusInternalServerError)
	}
}
