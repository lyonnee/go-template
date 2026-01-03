package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lyonnee/go-template/internal/interfaces/http/controller"
	"github.com/lyonnee/go-template/internal/interfaces/http/middleware"
	"github.com/lyonnee/go-template/pkg/di"
	"github.com/lyonnee/go-template/pkg/log"
)

func RegisterRoutes(engine *gin.Engine) {
	logger := di.Get[*log.Logger]()

	// register middleware
	engine.Use(middleware.Logger(logger))
	engine.Use(middleware.Recovery())
	engine.Use(middleware.CORS())
	engine.Use(middleware.AddTrace())

	// register handler
	apiRouter := engine.Group("/api")

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
