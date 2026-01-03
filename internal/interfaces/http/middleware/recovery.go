package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lyonnee/go-template/pkg/di"
	"github.com/lyonnee/go-template/pkg/log"
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

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		getRecoveryLogger().Errorf("[Recovery] err=%v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	})
}
