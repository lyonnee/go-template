package bootstrap

import (
	"github.com/lyonnee/go-template/internal/application/command_executor"
	"github.com/lyonnee/go-template/internal/application/query_executor"
	"github.com/lyonnee/go-template/internal/bootstrap/server"
	"github.com/lyonnee/go-template/internal/infrastructure/repository_impl"
	"github.com/lyonnee/go-template/internal/interfaces/http/controller"
	"github.com/lyonnee/go-template/pkg/container"
)

// Run initializes the bootstrap process for the application.
func Run() {
	registerServices()

	go server.StartHTTPServer()
	go server.StartRPCServer()
}

func registerServices() error {
	// controller
	container.AddSingletonService(controller.NewAuthController)
	container.AddSingletonService(controller.NewUserController)
	container.AddSingletonService(controller.NewHealthController)

	// application service
	container.AddSingletonService(command_executor.NewAuthCommandService)
	container.AddSingletonService(command_executor.NewUserCommandService)

	container.AddSingletonService(query_executor.NewUserQueryService)

	// repository
	container.AddTransientService(repository_impl.NewUserRepository)

	return nil
}
