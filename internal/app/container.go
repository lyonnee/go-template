package app

import (
	"github.com/lyonnee/go-template/internal/application/command_executor"
	"github.com/lyonnee/go-template/internal/application/query_executor"
	"github.com/lyonnee/go-template/internal/domain/repository"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	repoImpl "github.com/lyonnee/go-template/internal/infrastructure/repository"
	"github.com/lyonnee/go-template/internal/interfaces/http/controller"
)

// Container 依赖注入容器
type Container struct {
	// 通用组件
	logger log.Logger

	// Repository 层
	userRepository repository.UserRepository

	// Application Command Service 层
	authCmdService *command_executor.AuthCommandService
	userCmdService *command_executor.UserCommandService

	// Application Query Service 层
	userQueryService *query_executor.UserQueryService

	// Controller 层
	authController   *controller.AuthController
	userController   *controller.UserController
	healthController *controller.HealthController
}

// NewContainer 创建新的容器实例
func NewContainer() *Container {
	container := &Container{}

	container.initializeCommonComponents()
	container.initializeRepositories()
	container.initializeApplicationServices()
	container.initializeControllers()
	return container
}

// initializeCommonComponents 初始化通用组件
func (c *Container) initializeCommonComponents() {
	// 临时使用 noop logger，实际使用时应该注入真实的 logger
	c.logger = log.NewNoopLogger()
}

// initializeRepositories 初始化仓储层
func (c *Container) initializeRepositories() {
	c.userRepository = repoImpl.NewUserRepository(c.logger)
}

// initializeApplicationServices 初始化应用服务层
func (c *Container) initializeApplicationServices() {
	// Command Service
	c.authCmdService = command_executor.NewAuthCommandService(c.userRepository, c.logger)
	c.userCmdService = command_executor.NewUserCommandService(c.userRepository, c.logger)

	// Query Service
	c.userQueryService = query_executor.NewUserQueryService(c.userRepository, c.logger)
}

// initializeControllers 初始化控制器层
func (c *Container) initializeControllers() {
	c.authController = controller.NewAuthController(c.authCmdService, c.logger)
	c.userController = controller.NewUserController(c.userCmdService, c.userQueryService, c.logger)
	c.healthController = controller.NewHealthController()
}

// Logger 获取日志记录器
func (c *Container) Logger() log.Logger {
	return c.logger
}

// AuthController 获取认证控制器
func (c *Container) AuthController() *controller.AuthController {
	return c.authController
}

// UserController 获取用户控制器
func (c *Container) UserController() *controller.UserController {
	return c.userController
}

// HealthController 获取健康检查控制器
func (c *Container) HealthController() *controller.HealthController {
	return c.healthController
}
