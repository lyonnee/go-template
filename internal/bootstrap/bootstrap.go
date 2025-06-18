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
	di.AddSingletonController(controller.NewAuthController)
	di.AddSingletonController(controller.NewUserController)
	di.AddSingletonController(controller.NewHealthController)

	// application service
	// command
	di.AddSingletonService(command_executor.NewAuthCommandService)
	di.AddSingletonService(command_executor.NewUserCommandService)
	// query
	di.AddSingletonService(query_executor.NewUserQueryService)

	// repository
	di.AddTransientRepository(repository_impl.NewUserRepository)

	return nil
}
