package http

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/lyonnee/go-template/internal/application/services"
	"github.com/lyonnee/go-template/internal/interfaces/http/controller"
)

func New() *fiber.App {
	router := fiber.New()

	// register
	api := router.Group("/api")

	addV1(api)

	return router
}

func addV1(r fiber.Router) {
	// auth
	{
		authController := controller.NewAuthController(
			services.NewAuthService(),
		)

		authRouter := r.Group("/auth")
		authRouter.Post("/login", authController.Login)
		authRouter.Post("/refresh", authController.RefreshToken)
	}
	// other
}
