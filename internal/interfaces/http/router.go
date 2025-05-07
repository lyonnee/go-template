package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lyonnee/go-template/internal/application/service"
	"github.com/lyonnee/go-template/internal/interfaces/http/controller"
	"github.com/lyonnee/go-template/internal/interfaces/http/middleware"
)

func New() *fiber.App {
	app := fiber.New()

	app.Use(middleware.Logger())

	// register
	router := app.Group("/api")

	addV1(router)

	return app
}

func addV1(r fiber.Router) {
	base := r.Group("v1")
	// auth
	{
		authController := controller.NewAuthController(
			service.NewAuthService(),
		)

		authRouter := base.Group("/auth")
		authRouter.Post("/login", authController.Login)
		authRouter.Post("/refresh", authController.RefreshToken)
	}
	// other
}
