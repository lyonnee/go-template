package middleware

import (
	"context"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/lyonnee/go-template/infrastructure/di"
	"github.com/lyonnee/go-template/infrastructure/log"
	"go.uber.org/zap"
)

var once sync.Once
var recoveryLogger *zap.SugaredLogger

func getRecoveryLogger() *zap.SugaredLogger {
	if recoveryLogger == nil {
		once.Do(func() {
			recoveryLogger = di.Get[*log.Logger]().Sugar()
		})
	}
	return recoveryLogger
}

func Recovery(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
	getRecoveryLogger().Errorf("[Recovery] err=%v\nstack=%s", err, stack)
	c.AbortWithStatusJSON(consts.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
}
