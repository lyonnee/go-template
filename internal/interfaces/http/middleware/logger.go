package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lyonnee/go-template/pkg/log"
	"go.uber.org/zap"
)

func Logger() fiber.Handler {
	// return fiberzap.New(fiberzap.Config{
	// 	Logger: log.Logger(),
	// 	Fields: []string{
	// 		"status", "method", "path", "query", "ip", "ua", "errors", "latency",
	// 	},
	// })
	return func(c *fiber.Ctx) error {
		start := time.Now() // 请求的时间

		var errStr string
		if chainErr := c.Next(); chainErr != nil { // 执行后续中间件
			errStr = chainErr.Error()
		}

		cost := time.Since(start)
		log.Info("",
			zap.Int("status", c.Response().StatusCode()),                // 状态码
			zap.String("method", c.Method()),                            // 请求的方法
			zap.String("path", c.Path()),                                // 请求的路径
			zap.String("query", c.Request().URI().QueryArgs().String()), // 请求的参数
			zap.String("ip", c.IP()),                                    // 请求的IP
			zap.String("user-agent", c.Get(fiber.HeaderUserAgent)),      // 请求头
			zap.String("errors", errStr),                                // 错误信息
			zap.Duration("cost", cost),                                  // 请求时间
		)

		return nil
	}
}
