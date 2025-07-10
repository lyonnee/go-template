package http

import (
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/lyonnee/go-template/infrastructure/di"
	"github.com/lyonnee/go-template/infrastructure/log"
	"github.com/lyonnee/go-template/interfaces/http/controller"
	"github.com/lyonnee/go-template/interfaces/http/middleware"
)

func registerRoutes(hz *server.Hertz) {
	logger := di.Get[*log.Logger]()

	// register middleware
	hz.Use(middleware.Logger(logger))
	hz.Use(recovery.Recovery(recovery.WithRecoveryHandler(middleware.Recovery)))
	hz.Use(middleware.CORS())
	hz.Use(middleware.AddTrace())

	// register handler
	apiRouter := hz.Group("/api")

	// 健康检查
	{
		healthController := di.Get[*controller.HealthController]()

		apiRouter.GET("/health", healthController.HealthCheck)
		apiRouter.GET("/ready", healthController.ReadinessCheck)
		apiRouter.GET("/live", healthController.LivenessCheck)
	}

	// 认证相关
	{
		authController := di.Get[*controller.AuthController]()

		authRouter := apiRouter.Group("/auth")
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/refresh", authController.RefreshToken)
	}

	// 用户相关 (需要认证)
	{
		userController := di.Get[*controller.UserController]()

		userRouter := apiRouter.Group("/users")
		userRouter.POST("", userController.Register)

		userRouter.Use(middleware.JWTAuth())
		userRouter.GET("/:id", userController.GetUser)
		userRouter.PUT("/:id/username", userController.UpdateUsername)
	}
}
