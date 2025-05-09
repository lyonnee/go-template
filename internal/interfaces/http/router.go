package http

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/lyonnee/go-template/internal/application/service"
	"github.com/lyonnee/go-template/internal/interfaces/http/controller"
	"github.com/lyonnee/go-template/internal/interfaces/http/middleware"
)

func Register(hz *server.Hertz) {
	// process panic
	hz.PanicHandler = panicHandler

	// register middleware
	hz.Use(middleware.AddTrace())
	hz.Use(middleware.Logger())

	// register handler
	apiRouter := hz.Group("/api")

	addV1(apiRouter)
}

func panicHandler(c context.Context, ctx *app.RequestContext) {
}

func addV1(r *route.RouterGroup) {
	base := r.Group("v1")
	// auth
	{
		authController := controller.NewAuthController(
			service.NewAuthService(),
		)

		authRouter := base.Group("/auth")
		authRouter.POST("/user", authController.SignUp)
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/refresh", authController.RefreshToken)
	}
	// other
}
