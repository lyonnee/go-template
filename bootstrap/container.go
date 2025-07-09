package bootstrap

import (
	appService "github.com/lyonnee/go-template/application/service"
	"github.com/lyonnee/go-template/bootstrap/di"
	"github.com/lyonnee/go-template/infrastructure/repository_impl"
	"github.com/lyonnee/go-template/interfaces/http/controller"

	domainService "github.com/lyonnee/go-template/domain/service"
)

func registerServices() error {
	// controller
	di.AddSingleton(controller.NewAuthController)
	di.AddSingleton(controller.NewUserController)
	di.AddSingleton(controller.NewHealthController)

	// application service
	di.AddSingleton(appService.NewAuthCommandService)
	di.AddSingleton(appService.NewUserCommandService)

	di.AddSingleton(appService.NewUserQueryService)

	// domain service
	di.AddSingleton(domainService.NewUserService)

	// repository
	di.AddSingleton(repository_impl.NewUserRepository)

	return nil
}
