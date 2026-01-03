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
go run . -env dev
```

## Template Features

This template provides a production-ready Go web service with:

### 🏗️ Architecture Design
- 📦 **Clean Architecture** - Strict adherence to Clean Architecture principles with clear layering
- 🎯 **Domain-Driven Design** - DDD practices with rich domain models and business logic encapsulation
- 🔌 **CQRS Pattern** - Command Query Responsibility Segregation for read/write model separation
- 🏗️ **Dependency Injection** - IoC container based on samber/do/v2

### 🚀 Technology Stack
- ⚡ **High-Performance HTTP** - Gin web framework
- 🗄️ **Database Support** - PostgreSQL + SQLx with transaction management
- 💾 **Caching Layer** - Redis integration with context management
- 🔐 **Authentication System** - Complete JWT + OAuth implementation
- 📝 **Structured Logging** - Zap + Lumberjack logging system
- ⏰ **Task Scheduling** - Cron job support
- 🆔 **ID Generation** - Snowflake distributed ID generation

### 🔧 Development Support
- 🔧 **Multi-Environment Config** - dev/test/prod environment configuration management
- 🐳 **Docker Ready** - Multi-stage build Dockerfile
- 🔄 **Graceful Shutdown** - Complete resource cleanup mechanism
- 📊 **Health Checks** - Complete health check endpoints
- 🛡️ **Middleware System** - CORS, Recovery, JWT and other middleware
- 🔒 **Password Security** - bcrypt password hashing storage

## Project Structure

This template strictly follows Clean Architecture and Domain-Driven Design (DDD) principles:

```
go-template/
├── main.go                           # Program entrypoint
├── go.mod                            # Go module definition
├── go.sum                            # Dependency checksums
├── Dockerfile                        # Container deployment
├── _logs/                            # Local logs directory
│   └── dev.log
│
├── configs/                          # Multi-environment configuration
│   ├── config.dev.yaml              # Development environment config
│   ├── config.test.yaml              # Test environment config
│   └── config.prod.yaml              # Production environment config
│
├── internal/                         # Core business code (Go internal package)
│   ├── application/                  # Application layer - use case orchestration
│   │   ├── commands/                 # CQRS write side (commands)
│   │   │   ├── auth_command_service.go
│   │   │   └── user_command_service.go
│   │   ├── queries/                  # CQRS read side (queries)
│   │   │   └── user_query_service.go
│   │   └── scheduler/                # Task scheduling
│   │       ├── scheduler.go
│   │       └── jobs/
│   │           └── test_job.go
│   │
│   ├── domain/                       # Domain layer - business logic core
│   │   ├── entity/                   # Domain entities
│   │   │   └── user.go
│   │   ├── errors/                   # Domain error definitions
│   │   │   └── user_errors.go
│   │   ├── repository/               # Repository interfaces (ports)
│   │   │   ├── repository.go
│   │   │   ├── user_repository.go
│   │   │   └── eth_repository.go
│   │   └── service/                  # Domain services
│   │       ├── user_service.go
│   │       └── infra_service.go
│   │
│   ├── infrastructure/               # Infrastructure layer - technical implementations
│   │   ├── config/                   # Configuration management
│   │   │   ├── config.go
│   │   │   └── types.go
│   │   ├── database/                 # Database infrastructure
│   │   │   ├── database.go
│   │   │   ├── dbcontext.go
│   │   │   ├── logger.go
│   │   │   └── postgres.go
│   │   ├── cache/                    # Cache infrastructure
│   │   │   ├── cache.go
│   │   │   ├── cachecontext.go
│   │   │   ├── keys.go
│   │   │   └── redis.go
│   │   ├── auth/                     # Authentication infrastructure
│   │   │   ├── auth.go
│   │   │   ├── jwt.go
│   │   │   └── oauth.go
│   │   ├── repository_impl/          # Repository implementations (adapters)
│   │   │   ├── user_repository.go
│   │   │   └── model/                # Database models
│   │   │       ├── base_model.go
│   │   │       └── user.go
│   │   ├── blockchain/               # Blockchain utilities
│   │   │   └── blockchain.go
│   │   └── mq/                       # Message queue
│   │       └── mq.go
│   │
│   └── interfaces/                   # Interface adaptation layer - external interfaces
│       ├── http/                     # HTTP interfaces
│       │   ├── router.go             # Route configuration
│       │   ├── controller/           # HTTP controllers
│       │   │   ├── auth_controller.go
│       │   │   ├── user_controller.go
│       │   │   └── health_controller.go
│       │   ├── dto/                  # Data transfer objects
│       │   │   ├── base_response.go
│       │   │   ├── auth.go
│       │   │   ├── user.go
│       │   │   └── pagequery.go
│       │   └── middleware/           # HTTP middleware
│       │       ├── cors.go
│       │       ├── jwt.go
│       │       ├── logger.go
│       │       ├── recovery.go
│       │       └── trace.go
│       ├── grpc/                     # gRPC interfaces
│       └── event_handler/            # Event handlers
│           └── event_handler.go
│
├── pkg/                              # Shared utility libraries
│   ├── di/                           # Dependency injection container
│   │   └── injector.go
│   ├── log/                          # Logging utilities
│   │   ├── log.go
│   │   └── zap_logger.go
│   ├── idgen/                        # ID generator
│   │   └── id_generator.go
│   └── util/                         # General utilities
│       └── bcrypt.go
│
├── services/                         # Service startup management
│   ├── services.go                   # Service registry
│   ├── http.go                       # HTTP service
│   ├── grpc.go                       # gRPC service
│   └── cron.go                       # Scheduled task service
│
├── scripts/                          # Build scripts
│   ├── build.sh
│   └── start.sh
│
├── sqls/                             # Database migrations
│   └── user.sql
│
└── test/                             # Test directory
```

### Architecture Layer Description

#### 1. Interface Adaptation Layer (Interfaces Layer)
- **HTTP Controllers**: Handle RESTful API request/response, responsible for parameter binding and response formatting
- **Middleware System**: Cross-cutting concerns implementation (authentication, logging, CORS, recovery, etc.)
- **DTO Objects**: Data transfer objects for data exchange between interface layer and application layer
- **Route Management**: Unified route registration and management

#### 2. Application Layer (Application Layer)
- **Command Services**: CQRS write side, handling business operations and state changes
- **Query Services**: CQRS read side, handling data queries and display logic
- **Task Scheduling**: Scheduled tasks and background job orchestration management
- **Transaction Coordination**: Cross-domain object transaction management and data consistency

#### 3. Domain Layer (Domain Layer)
- **Domain Entities**: Core business objects with unique identity, containing business logic
- **Domain Services**: Business logic that doesn't fit well within a single entity
- **Repository Interfaces**: Data access abstractions (ports)
- **Domain Errors**: Business-related error definitions and handling

#### 4. Infrastructure Layer (Infrastructure Layer)
- **Repository Implementations**: Concrete data persistence implementations (adapters)
- **Database Access**: PostgreSQL connection, transaction management, query logging
- **Cache Implementation**: Redis caching strategies and context management
- **Authentication System**: JWT token generation/validation, OAuth integration
- **Configuration Management**: Multi-environment configuration loading and type-safe access

## Development Guide

### Adding New Business Features

Follow these steps to add new functionality while maintaining architectural integrity:

#### 1. Define Domain Entity
```go
// internal/domain/entity/product.go
type Product struct {
    ID          uint64    `json:"id"`
    Name        string    `json:"name"`
    Price       float64   `json:"price"`
    CreatedAt   int64     `json:"created_at"`
    UpdatedAt   int64     `json:"updated_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
    // Validation and creation logic
}
```

#### 2. Create Repository Interface
```go
// internal/domain/repository/product_repository.go
type ProductRepository interface {
    BaseRepository
    Save(ctx context.Context, product *entity.Product) error
    FindByID(ctx context.Context, id uint64) (*entity.Product, error)
    FindAll(ctx context.Context) ([]*entity.Product, error)
    Delete(ctx context.Context, id uint64) error
}
```

#### 3. Implement Repository
```go
// internal/infrastructure/repository_impl/product_repository.go
type ProductRepoImpl struct {
    BaseRepoImpl
}

func init() {
    di.AddSingleton[repository.ProductRepository](NewProductRepository)
}

func NewProductRepository() (repository.ProductRepository, error) {
    return &ProductRepoImpl{}, nil
}
```

#### 4. Create Application Services
```go
// internal/application/commands/product_command_service.go
type ProductCommandService struct {
    productRepo repository.ProductRepository
    logger      *log.Logger
}

func init() {
    di.AddSingleton[*ProductCommandService](NewProductCommandService)
}
```

#### 5. Add HTTP Controller
```go
// internal/interfaces/http/controller/product_controller.go
type ProductController struct {
    productCmdService *commands.ProductCommandService
    logger           *log.Logger
}

func init() {
    di.AddSingleton[*ProductController](NewProductController)
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
    // HTTP request handling
}
```

#### 6. Register Routes
Update `internal/interfaces/http/router.go`:
```go
// Product routes
productController := di.Get[*ProductController]()
apiRouter.POST("/products", productController.CreateProduct)
apiRouter.GET("/products/:id", productController.GetProduct)
```

### Dependency Injection Usage Guide

This project uses a custom dependency injection system based on samber/do/v2. **All services (repositories, domain services, application services, controllers) must be registered through the unified DI interface.**

#### Core Principles

1. **Self-Registration**: Services register themselves using `init()` functions
2. **Factory Pattern**: Use `New` functions as service factories  
3. **Type Safety**: Leverage Go generics for type safety
4. **Unified Interface**: All dependency registrations use `di.AddSingleton` or `di.AddTransient`

#### Registration Examples

**Repository Registration:**
```go
func init() {
    di.AddSingleton[repository.UserRepository](NewUserRepository)
}
```

**Service Registration:**
```go
func init() {
    di.AddSingleton[*UserService](NewUserService)
}
```

**Controller Registration:**
```go
func init() {
    di.AddSingleton[*UserController](NewUserController)
}
```

#### Best Practices
- Register interfaces for repositories and domain services
- Use concrete types for application services and controllers
- Always provide factory functions that return `(T, error)`
- Handle dependency resolution in factory functions using `di.Get[T]()`

## API Documentation

### Health Check Endpoints
```
GET /api/health      # Health status check
GET /api/ready       # Readiness check  
GET /api/live        # Liveness check
```

### Authentication Endpoints
```
POST /api/auth/login    # User login
POST /api/auth/refresh  # Refresh token
```

### User Endpoints
```
POST /api/users             # User registration
GET  /api/users/:id         # Get user info (requires auth)
PUT  /api/users/:id/username # Update username (requires auth)
```

## System Requirements

- **Go**: 1.23.7 or higher
- **PostgreSQL**: 12 or higher
- **Redis**: 6.0 or higher
- **Docker**: 20.10 or higher (optional)

## Configuration

The project supports multi-environment configuration via the `-env` parameter:

```bash
# Development environment (default)
go run . -env dev

# Test environment
go run . -env test  

# Production environment
go run . -env prod
```

Configuration files are located in the `configs/` directory:
- `config.dev.yaml` - Development environment configuration
- `config.test.yaml` - Test environment configuration
- `config.prod.yaml` - Production environment configuration

## Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/domain/...
```

## Building and Running

```bash
# Development
go run . -env dev

# Build for production
./scripts/build.sh

# Docker build
docker build -t your-app .
docker run -p 8080:8080 your-app
```

## Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Follow the project's architecture principles and coding standards
4. Write tests for new functionality
5. Ensure all tests pass
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Create a Pull Request

### Code Standards
- Follow Go official code style
- Use meaningful variable and function names
- Provide documentation comments for public interfaces
- Keep functions concise with single responsibility

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Maintainer

- [@lyonnee](https://github.com/lyonnee)