package bootstrap

import (
	"github.com/lyonnee/go-template/application/service"
	"github.com/lyonnee/go-template/bootstrap/di"
	"github.com/lyonnee/go-template/infrastructure/repository_impl"
	"github.com/lyonnee/go-template/interfaces/http/controller"
)

func registerServices() error {
	// controller
	di.AddSingleton(controller.NewAuthController)
	di.AddSingleton(controller.NewUserController)
	di.AddSingleton(controller.NewHealthController)

	// application service
	di.AddSingleton(service.NewAuthCommandService)
	di.AddSingleton(service.NewUserCommandService)

	di.AddSingleton(service.NewUserQueryService)

	// repository
	di.AddTransient(repository_impl.NewUserRepository)

	return nil
}
