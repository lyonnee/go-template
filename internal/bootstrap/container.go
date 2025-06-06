package bootstrap

import (
	"github.com/lyonnee/go-template/internal/application/command_executor"
	"github.com/lyonnee/go-template/internal/application/query_executor"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/internal/infrastructure/repository_impl"
	"github.com/lyonnee/go-template/internal/interfaces/http/controller"
	"github.com/lyonnee/go-template/pkg/container"
)

func registerServices() error {
	// controller
	container.AddSingletonService[*controller.AuthController](controller.NewAuthController)
	container.AddSingletonService[*controller.UserController](controller.NewUserController)
	container.AddSingletonService[*controller.HealthController](controller.NewHealthController)

	// application service
	container.AddSingletonService[*command_executor.AuthCommandService](command_executor.NewAuthCommandService)
	container.AddSingletonService[*command_executor.UserCommandService](command_executor.NewUserCommandService)

	container.AddSingletonService[*query_executor.UserQueryService](query_executor.NewUserQueryService)

	// repository
	container.AddTransientService[repository.UserRepository](repository_impl.NewUserRepository)

	return nil
}
