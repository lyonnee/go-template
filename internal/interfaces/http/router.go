package http

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/lyonnee/go-template/internal/infrastructure/di"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/lyonnee/go-template/internal/interfaces/http/controller"
	"github.com/lyonnee/go-template/internal/interfaces/http/middleware"
)

func Register(hz *server.Hertz) {
	logger := di.GetService[log.Logger]()

	// process panic
	hz.PanicHandler = panicHandler

	// register middleware
	hz.Use(middleware.Logger(logger))
	hz.Use(recovery.Recovery(recovery.WithRecoveryHandler(middleware.Recovery(logger))))
	hz.Use(middleware.CORS())
	hz.Use(middleware.AddTrace())

	// register handler
	apiRouter := hz.Group("/api")

	addV1(apiRouter)
}

func panicHandler(c context.Context, ctx *app.RequestContext) {

}

func addV1(r *route.RouterGroup) {
	base := r.Group("v1")

	// 健康检查
	{
		healthController := di.GetService[*controller.HealthController]()

		base.GET("/health", healthController.HealthCheck)
		base.GET("/ready", healthController.ReadinessCheck)
		base.GET("/live", healthController.LivenessCheck)
	}

	// 认证相关
	{
		authController := di.GetService[*controller.AuthController]()

		authRouter := base.Group("/auth")
		authRouter.POST("/signup", authController.SignUp)
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/refresh", authController.RefreshToken)
	}

	// 用户相关 (需要认证)
	{
		userController := di.GetService[*controller.UserController]()

		userRouter := base.Group("/users")
		userRouter.Use(middleware.JWTAuth())
		userRouter.GET("/:id", userController.GetUser)
		userRouter.PUT("/:id/username", userController.UpdateUsername)
	}
}
