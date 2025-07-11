<div align="center">

# Go Template

| English | [中文](README_zh.md) |
| --- | --- |

A modern Go web service project template built with Domain-Driven Design (DDD) and Clean Architecture principles.
</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/lyonnee/go-template)](https://goreportcard.com/report/github.com/lyonnee/go-template)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lyonnee/go-template)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Quick Start

### Create Project from Template

1. Install gonew tool
```bash
go install golang.org/x/tools/cmd/gonew@latest
```

2. Create new project from template
```bash
gonew github.com/lyonnee/go-template github.com/your-username/your-project
```

3. Navigate to project directory
```bash
cd your-project
```

4. Start development server
```bash
go run cmd/server/main.go -e dev
```

## Template Features

This template provides a production-ready Go web service with:

- 📦 **Clean Architecture** - Well-structured layers with clear separation of concerns
- 🎯 **Domain-Driven Design** - Rich domain models and business logic encapsulation
- 🔐 **Authentication System** - JWT and OAuth integration ready
- 📝 **Logging Infrastructure** - Structured logging with multiple outputs
- 🗄️ **Database Support** - Multi-database compatibility
- 💾 **Caching Layer** - Redis integration
- 🔄 **Graceful Shutdown** - Proper resource cleanup
- 🐳 **Docker Ready** - Multi-stage builds included
- ⚡ **High-Performance HTTP** - CloudWeGo Hertz framework
- 🔌 **gRPC Support** - Protocol buffer integration
- 📨 **Message Queue** - Async processing infrastructure
- ⏰ **Task Scheduling** - Cron job support
- 🔧 **Multi-Environment Config** - Development, test, production configs
- 🏗️ **Dependency Injection** - IoC container included
- 🆔 **ID Generation** - Distributed ID generation
- 🧪 **Testing Structure** - Test organization and utilities

## Project Structure

This template follows Clean Architecture and DDD principles with the following structure:

```
go-template/
├── cmd/                             # Application entry points
│   ├── scheduler/                   # Background task scheduler
│   └── server/                      # Main HTTP/gRPC server
│
├── application/                     # Application Layer (Use Cases)
│   ├── cron/                        # Scheduled job definitions
│   └── service/                     # Application services
│       ├── auth_command_service.go  # Authentication operations
│       ├── user_command_service.go  # User write operations
│       └── user_query_service.go    # User read operations
│
├── domain/                          # Domain Layer (Business Logic)
│   ├── entity/                      # Business entities
│   ├── errors/                      # Domain-specific errors
│   ├── repository/                  # Repository interfaces
│   └── service/                     # Domain services
│
├── infrastructure/                  # Infrastructure Layer (Technical Details)
│   ├── auth/                        # Authentication implementations
│   ├── blockchain/                  # Blockchain integrations
│   ├── cache/                       # Cache implementations
│   ├── config/                      # Configuration management
│   ├── database/                    # Database connections
│   ├── di/                          # Dependency injection container
│   ├── log/                         # Logging implementations
│   ├── mq/                          # Message queue implementations
│   └── repository_impl/             # Repository implementations
│       └── model/                   # Database models
│
├── interfaces/                      # Interface Layer (External Interface)
│   ├── event_handler/               # Event handlers
│   ├── grpc/                        # gRPC service definitions
│   └── http/                        # HTTP interface
│       ├── controller/              # HTTP request handlers
│       ├── dto/                     # Data transfer objects
│       ├── middleware/              # HTTP middlewares
│       ├── router.go                # Route definitions
│       └── server.go                # HTTP server setup
│
├── pkg/                             # Shared utilities
│   └── idgen/                       # ID generation utilities
│
├── scripts/                         # Build and deployment scripts
├── sqls/                            # Database schema files
├── test/                            # Test files and utilities
│
├── config.dev.yaml                  # Development environment config
├── config.test.yaml                 # Test environment config
├── config.prod.yaml                 # Production environment config
├── Dockerfile                       # Container definition
└── docker-compose.yml               # Multi-service setup
```

## Architecture Overview

This project implements **Clean Architecture** with **Domain-Driven Design (DDD)** principles:

### 1. **Interface Layer** (External Interface)
- **HTTP Controllers**: Handle REST API requests
- **gRPC Services**: Handle RPC requests
- **Middlewares**: Cross-cutting concerns (CORS, logging, authentication)
- **DTOs**: Data transfer objects for external communication

### 2. **Application Layer** (Use Cases)
- **Command Services**: Handle write operations
- **Query Services**: Handle read operations
- **Application Services**: Orchestrate business workflows

### 3. **Domain Layer** (Business Logic)
- **Entities**: Core business objects with identity
- **Domain Services**: Business logic that spans multiple entities
- **Repository Interfaces**: Data access abstractions
- **Domain Events**: Business event definitions

### 4. **Infrastructure Layer** (Technical Details)
- **Repository Implementations**: Data persistence implementations
- **Cache Implementations**: Caching strategies
- **Message Queue**: Async communication
- **Configuration**: Environment-specific settings

## Development Guide

### Adding New Business Features

#### 1. Define Domain Entity
Create new business entities in `domain/entity/`:

```go
// domain/entity/product.go
type Product struct {
    ID          int64     `json:"id"`
    Name        string    `json:"name"`
    Price       float64   `json:"price"`
    CreatedAt   int64     `json:"created_at"`
    UpdatedAt   int64     `json:"updated_at"`
}
```

#### 2. Create Repository Interface
Define data access interface in `domain/repository/`:

```go
// domain/repository/product_repository.go
type ProductRepository interface {
    Save(ctx context.Context, product *entity.Product) error
    FindByID(ctx context.Context, id int64) (*entity.Product, error)
    FindAll(ctx context.Context) ([]*entity.Product, error)
    Delete(ctx context.Context, id int64) error
}
```

#### 3. Implement Repository
Create concrete implementation in `infrastructure/repository_impl/`:

```go
// infrastructure/repository_impl/product_repo_impl.go
type ProductRepoImpl struct {
    db *sqlx.DB
}

func (r *ProductRepoImpl) Save(ctx context.Context, product *entity.Product) error {
    // Database implementation
}
```

#### 4. Register Repository to DI Container
Register the repository in the same file using `init()` function:

```go
// infrastructure/repository_impl/product_repo_impl.go
type ProductRepoImpl struct {
    db *sqlx.DB
}

func init() {
    di.AddSingleton[repository.ProductRepository](NewProductRepository)
}

func NewProductRepository() (repository.ProductRepository, error) {
    db := di.Get[*sqlx.DB]()
    return &ProductRepoImpl{db: db}, nil
}

func (r *ProductRepoImpl) Save(ctx context.Context, product *entity.Product) error {
    // Database implementation
}
```

#### 4. Create Application Service
Implement business logic in `application/service/`:

```go
// application/service/product_service.go
type ProductService struct {
    productRepo repository.ProductRepository
}

func (s *ProductService) CreateProduct(ctx context.Context, req CreateProductRequest) error {
    // Business logic implementation
}
```

#### 5. Register Application Service to DI Container
Register the application service in the same file:

```go
// application/service/product_service.go
type ProductService struct {
    productRepo repository.ProductRepository
}

func init() {
    di.AddSingleton[*ProductService](NewProductService)
}

func NewProductService() (*ProductService, error) {
    repo := di.Get[repository.ProductRepository]()
    return &ProductService{productRepo: repo}, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, req CreateProductRequest) error {
    // Business logic implementation
}
```

#### 5. Add HTTP Controller
Handle HTTP requests in `interfaces/http/controller/`:

```go
// interfaces/http/controller/product_controller.go
type ProductController struct {
    productService *service.ProductService
}

func (c *ProductController) CreateProduct(ctx context.Context, req *app.RequestContext) {
    // HTTP request handling
}
```

#### 6. Register Controller to DI Container
Register the controller in the same file:

```go
// interfaces/http/controller/product_controller.go
type ProductController struct {
    productService *service.ProductService
}

func init() {
    di.AddSingleton[*ProductController](NewProductController)
}

func NewProductController() (*ProductController, error) {
    service := di.Get[*ProductService]()
    return &ProductController{productService: service}, nil
}

func (c *ProductController) CreateProduct(ctx context.Context, req *app.RequestContext) {
    // HTTP request handling
}
```

#### 6. Register Routes
Update routes in `interfaces/http/router.go`:

```go
// Add to router.go
productController := di.Get[*ProductController]()
v1.POST("/products", productController.CreateProduct)
v1.GET("/products/:id", productController.GetProduct)
v1.PUT("/products/:id", productController.UpdateProduct)
v1.DELETE("/products/:id", productController.DeleteProduct)
```

### Adding New Configuration

#### 1. Update Configuration Structure
Add new config section in `infrastructure/config/types.go`:

```go
type Config struct {
    // ... existing fields
    Email EmailConfig `mapstructure:"email"`
}

type EmailConfig struct {
    Provider  string `mapstructure:"provider"`
    SMTPHost  string `mapstructure:"smtp_host"`
    SMTPPort  int    `mapstructure:"smtp_port"`
    Username  string `mapstructure:"username"`
    Password  string `mapstructure:"password"`
}
```

#### 2. Update Configuration Files
Add configuration to environment files:

```yaml
# config.dev.yaml
email:
  provider: smtp
  smtp_host: smtp.gmail.com
  smtp_port: 587
  username: your-email@gmail.com
  password: your-password
```

#### 3. Register Service
Use unified dependency injection interface:

```go
// Register during initialization
di.AddSingleton(func() (EmailService, error) {
    config := di.Get[*config.Config]()
    return &emailServiceImpl{
        config: config.Email,
    }, nil
})
```

### Dependency Injection Usage Guide

This template uses a unified dependency injection interface based on samber/do/v2. **All services (repositories, domain services, application services, controllers) must be registered through the interfaces provided by `injector.go`.**

#### Core Principles

1. **Self-Registration**: Services register themselves using `init()` functions
2. **Factory Pattern**: Use `New` functions as service factories
3. **Type Safety**: Leverage Go generics for type safety
4. **Unified Interface**: All dependency registrations use `di.AddSingleton` or `di.AddTransient`

#### Service Registration Pattern

Each service should follow this pattern:

```go
// domain/service/user_service.go
type UserService struct {
    logger   *log.Logger
    userRepo repository.UserRepository
}

func init() {
    di.AddSingleton[*UserService](NewUserService)
}

func NewUserService() (*UserService, error) {
    return &UserService{
        logger:   di.Get[*log.Logger](),
        userRepo: di.Get[repository.UserRepository](),
    }, nil
}

func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
    // Business logic implementation
}
```

#### Repository Registration

```go
// infrastructure/repository_impl/user_repo_impl.go
type UserRepoImpl struct {
    db *sqlx.DB
}

func init() {
    di.AddSingleton[repository.UserRepository](NewUserRepository)
}

func NewUserRepository() (repository.UserRepository, error) {
    db := di.Get[*sqlx.DB]()
    return &UserRepoImpl{db: db}, nil
}
```

#### Application Service Registration

```go
// application/service/user_command_service.go
type UserCommandService struct {
    userRepo      repository.UserRepository
    userDomainSvc *domain.UserService
}

func init() {
    di.AddSingleton[*UserCommandService](NewUserCommandService)
}

func NewUserCommandService() (*UserCommandService, error) {
    return &UserCommandService{
        userRepo:      di.Get[repository.UserRepository](),
        userDomainSvc: di.Get[*domain.UserService](),
    }, nil
}
```

#### Controller Registration

```go
// interfaces/http/controller/user_controller.go
type UserController struct {
    userCommandService *service.UserCommandService
    userQueryService   *service.UserQueryService
}

func init() {
    di.AddSingleton[*UserController](NewUserController)
}

func NewUserController() (*UserController, error) {
    return &UserController{
        userCommandService: di.Get[*service.UserCommandService](),
        userQueryService:   di.Get[*service.UserQueryService](),
    }, nil
}
```

#### Transient Service Registration

For services that need new instances each time:

```go
// infrastructure/email/email_service.go
type EmailService struct {
    config *config.EmailConfig
}

func init() {
    di.AddTransient[*EmailService](NewEmailService)
}

func NewEmailService() (*EmailService, error) {
    config := di.Get[*config.Config]()
    return &EmailService{config: &config.Email}, nil
}
```

#### Getting Services

```go
// Get services in routes or other components
func SetupRoutes() {
    userController := di.Get[*controller.UserController]()
    v1.POST("/users", userController.CreateUser)
}

// Get dependencies in services
func (s *SomeService) ProcessUser() {
    userRepo := di.Get[repository.UserRepository]()
    // Use repository...
}
```

#### Important Best Practices

- **Self-Registration**: Each service registers itself in its own `init()` function
- **Factory Functions**: Always provide a `New` function as the service factory
- **Interface Registration**: Prefer registering interface types for repositories and domain services
- **Concrete Registration**: Use concrete types for application services and controllers
- **Dependency Injection**: Always use `di.Get[T]()` to resolve dependencies in factory functions
- **Error Handling**: Factory functions should return `(T, error)` for proper error handling

### Adding New Middleware

#### 1. Create Middleware
Add new middleware in `interfaces/http/middleware/`:

```go
// interfaces/http/middleware/rate_limit.go
func RateLimit() app.HandlerFunc {
    return func(ctx context.Context, c *app.RequestContext) {
        // Rate limiting logic
        c.Next(ctx)
    }
}
```

#### 2. Register Middleware
Update router to use middleware:

```go
// interfaces/http/router.go
h.Use(middleware.RateLimit())
```

### Adding New Services

#### 1. Create Service Interface
Define service contract in `domain/service/`:

```go
// domain/service/notification_service.go
type NotificationService interface {
    SendEmail(ctx context.Context, to, subject, body string) error
    SendSMS(ctx context.Context, to, message string) error
}
```

#### 2. Implement Domain Service
Create implementation in `domain/service/`:

```go
// domain/service/notification_service_impl.go
type NotificationServiceImpl struct {
    config EmailConfig
}

func (s *NotificationServiceImpl) SendEmail(ctx context.Context, to, subject, body string) error {
    // Email sending implementation
}
```

#### 3. Register Domain Service to DI Container
Implement and register the domain service in the same file:

```go
// domain/service/notification_service_impl.go
type NotificationServiceImpl struct {
    config EmailConfig
    logger *log.Logger
}

func init() {
    di.AddSingleton[NotificationService](NewNotificationService)
}

func NewNotificationService() (NotificationService, error) {
    config := di.Get[*config.Config]()
    logger := di.Get[*log.Logger]()
    return &NotificationServiceImpl{
        config: config.Email,
        logger: logger,
    }, nil
}

func (s *NotificationServiceImpl) SendEmail(ctx context.Context, to, subject, body string) error {
    // Email sending implementation
}
```

### Adding Database Models

#### 1. Create Database Model
Add model in `infrastructure/repository_impl/model/`:

```go
// infrastructure/repository_impl/model/product.go
type Product struct {
    BaseModel
    Name  string  `db:"name"`
    Price float64 `db:"price"`
}
```

#### 2. Create Migration
Add SQL file in `sqls/`:

```sql
-- sqls/product.sql
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    deleted_at BIGINT NOT NULL DEFAULT 0,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL
);
```

### Testing Your Changes

```bash
# Run all tests
go test ./...

# Run specific package tests
go test ./domain/...
go test ./application/...

# Run with coverage
go test -cover ./...
```

### Building and Running

```bash
# Build application
./scripts/build.sh

# Run development server
go run cmd/server/main.go -e dev

# Run with Docker
docker build -t your-app .
docker run -p 8080:8080 your-app
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Follow the template structure and architecture principles
4. Write tests for your changes
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Maintainer

- [@lyonnee](https://github.com/lyonnee)