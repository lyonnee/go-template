package bootstrap

import (
	"github.com/lyonnee/go-template/internal/application/command_executor"
	"github.com/lyonnee/go-template/internal/application/query_executor"
	"github.com/lyonnee/go-template/internal/bootstrap/server"
	"github.com/lyonnee/go-template/internal/infrastructure/di"
	"github.com/lyonnee/go-template/internal/infrastructure/repository_impl"
	"github.com/lyonnee/go-template/internal/interfaces/http/controller"
)

func Initialize() {
	registerServices()
}

// Run initializes the bootstrap process for the application.
func Run() {
	go server.StartHTTPServer()
	go server.StartRPCServer()
}

func registerServices() error {
	// controller
	di.AddSingletonService(controller.NewAuthController)
	di.AddSingletonService(controller.NewUserController)
	di.AddSingletonService(controller.NewHealthController)

	// application service
	di.AddSingletonService(command_executor.NewAuthCommandService)
	di.AddSingletonService(command_executor.NewUserCommandService)

	di.AddSingletonService(query_executor.NewUserQueryService)

	// repository
	di.AddTransientService(repository_impl.NewUserRepository)

	return nil
}
