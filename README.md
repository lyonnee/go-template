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

### Using gonew

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
go run cmd/server/main.go -env dev
```

### Environment Configuration

Set the runtime environment through:
- Command line flags: `-env` (dev, test, prod)
- Environment variable: `APP_ENV`
- Default value: `dev`

## Features

- 📦 **Clean Architecture** - Well-structured layers with clear separation of concerns
- 🎯 **Domain-Driven Design (DDD)** - Rich domain models and business logic encapsulation
- 🔐 **JWT Authentication** - Secure token-based authentication with refresh tokens
- 📝 **Structured Logging** - Zap logger with configurable levels and file rotation
- 🗄️ **Multi-Database Support** - PostgreSQL and MySQL with connection pooling
- 💾 **Redis Caching** - Distributed caching with configurable prefix
- 🔄 **Graceful Shutdown** - Proper resource cleanup on termination
- 🐳 **Docker Support** - Multi-stage builds with optimized images
- ⚡ **High-Performance HTTP** - CloudWeGo Hertz framework
- 🔌 **gRPC Support** - Protocol buffer based RPC communication
- 📨 **Message Queue** - Async processing capabilities
- ⏰ **Task Scheduling** - Cron-based job scheduling
- 🔧 **Multi-Environment Config** - YAML-based configuration management
- 🏗️ **Dependency Injection** - Clean IoC container implementation
- 🧪 **Testing Support** - Comprehensive test utilities and mocks

## Project Structure

```
go-template/
├── cmd/                             # Application entry points
│   ├── migrate/                     # Database migration tool
│   ├── scheduler/                   # Task scheduler service
│   └── server/                      # Main HTTP/gRPC server
│
├── config/                          # Configuration management
│   ├── config.go                    # Config loading and validation
│   ├── auth.go                      # Authentication config
│   ├── cache.go                     # Cache config
│   ├── http.go                      # HTTP server config
│   ├── log.go                       # Logging config
│   └── persistence.go               # Database config
│
├── internal/                        # Private application code
│   ├── app/
│   │   └── container.go             # Dependency injection container
│   ├── application/                 # Application layer (Use Cases)
│   │   ├── command_service/        # Command handlers (CQRS)
│   │   │   ├── auth_command_service.go
│   │   │   └── user_command_service.go
│   │   └── query_executor/          # Query handlers (CQRS)
│   │       └── user_query_service.go
│   ├── domain/                      # Domain layer (Business Logic)
│   │   ├── entity/                  # Domain entities
│   │   │   └── user.go
│   │   ├── errors/                  # Domain-specific errors
│   │   │   └── user_errors.go
│   │   ├── repository/              # Repository interfaces
│   │   │   ├── repository.go
│   │   │   └── user_repository.go
│   │   ├── valueobject/             # Value objects
│   │   │   └── user_valueobjects.go
│   │   └── user_domain_service.go   # Domain services
│   ├── infrastructure/              # Infrastructure layer
│   │   ├── cache/                   # Cache implementations
│   │   │   ├── cache.go
│   │   │   └── keys.go
│   │   ├── eventbus/                # Event bus implementation
│   │   ├── log/                     # Logging implementations
│   │   │   ├── logger.go
│   │   │   └── noop_logger.go
│   │   ├── repository/              # Repository implementations
│   │   │   ├── model/               # Database models
│   │   │   ├── user_repo_impl.go
│   │   │   └── user_repo_impl_test.go
│   │   └── scheduler/               # Task scheduling
│   │       ├── cron.go
│   │       └── job_registry.go
│   └── interfaces/                  # Interface layer
│       ├── grpc/                    # gRPC interfaces
│       │   └── user.proto
│       └── http/                    # HTTP interfaces
│           ├── controller/          # HTTP controllers
│           ├── dto/                 # Data transfer objects
│           ├── middleware/          # HTTP middlewares
│           └── router.go            # Route definitions
│
├── pkg/                             # Public packages
│   ├── auth/                        # Authentication utilities
│   │   ├── jwt.go                   # JWT token handling
│   │   ├── oauth.go                 # OAuth integration
│   │   └── password.go              # Password hashing
│   ├── blockchain/                  # Blockchain integration
│   │   └── blockchain.go
│   ├── cache/                       # Cache utilities
│   │   └── cache.go
│   ├── hash/                        # Hashing utilities
│   ├── log/                         # Logging utilities
│   │   ├── zap_logger.go
│   │   └── zap_sugar_logger.go
│   ├── mq/                          # Message queue utilities
│   │   └── mq.go
│   └── persistence/                 # Database utilities
│       ├── persistence.go
│       └── postgres.go
│
├── server/                          # Server implementations
│   ├── http.go                      # HTTP server setup
│   └── rpc.go                       # gRPC server setup
│
├── scripts/                         # Build and deployment scripts
│   ├── build.sh                     # Build script
│   └── start.sh                     # Start script
│
├── sql/                             # Database schemas
│   └── user.sql                     # User table definitions
│
├── config.dev.yaml                  # Development configuration
├── config.test.yaml                 # Test configuration
├── config.prod.yaml                 # Production configuration
├── Dockerfile                       # Container definition
└── docker-compose.yml               # Multi-service setup
```

## Architecture Overview

This project implements **Clean Architecture** with **Domain-Driven Design (DDD)** principles:

### 1. **Interface Layer** (Adapters)
- **HTTP Controllers**: Handle REST API requests using Hertz framework
- **gRPC Services**: Handle RPC requests with Protocol Buffers
- **Middlewares**: Cross-cutting concerns (CORS, logging, authentication)
- **DTOs**: Data transfer objects for external communication

### 2. **Application Layer** (Use Cases)
- **Command Handlers**: Process write operations (CQRS pattern)
- **Query Handlers**: Process read operations (CQRS pattern)
- **Application Services**: Orchestrate business workflows
- **Transaction Management**: Ensure data consistency

### 3. **Domain Layer** (Business Logic)
- **Entities**: Core business objects with identity
- **Value Objects**: Immutable objects representing concepts
- **Domain Services**: Business logic that doesn't belong to entities
- **Repository Interfaces**: Data access abstractions
- **Domain Events**: Business event definitions

### 4. **Infrastructure Layer** (Frameworks & Drivers)
- **Repository Implementations**: Data persistence using SQLX
- **Cache Implementations**: Redis-based caching
- **Message Queue**: Async communication
- **Logging**: Structured logging with Zap
- **Scheduling**: Cron-based task execution

## Technology Stack

- **Language**: Go 1.23.7
- **HTTP Framework**: [CloudWeGo Hertz](https://github.com/cloudwego/hertz)
- **Database**: PostgreSQL, MySQL (via SQLX)
- **Cache**: Redis
- **Logging**: Uber Zap
- **Authentication**: JWT with refresh tokens
- **Configuration**: Viper (YAML)
- **Testing**: Testify, SQL Mock
- **Containerization**: Docker

## Configuration

The application uses YAML configuration files for different environments:

### Development (`config.dev.yaml`)
```yaml
http:
  port: :8081

log:
  enable_to_console: true
  to_console_level: debug
  to_file_level: debug
  filename: ./_logs/dev.log

persistence:
  postgres:
    dsn: postgres://postgres:admin123@localhost:5432/go_template?sslmode=disable
  mysql:
    dsn: root:admin123@tcp(localhost:3306)/go-template?charset=utf8mb4&parseTime=true

auth:
  jwt:
    secret_key: go-template
    access_token_expiry: 15m
    refresh_token_expiry: 168h

cache:
  redis:
    host: localhost
    port: 6379
    database: 0
    prefix: "go-template:"
```

## Development Guide

### Prerequisites
- Go 1.23.7 or later
- PostgreSQL or MySQL
- Redis (optional, for caching)
- Docker (optional, for containerization)

### Local Development

1. **Clone the repository**
```bash
git clone https://github.com/lyonnee/go-template.git
cd go-template
```

2. **Install dependencies**
```bash
go mod download
```

3. **Set up database**
```bash
# Run migrations
go run cmd/migrate/main.go
```

4. **Start the server**
```bash
# Development mode
go run cmd/server/main.go -env dev

# Or use the build script
./scripts/build.sh
./bin/server -env dev
```

5. **Start the scheduler (optional)**
```bash
go run cmd/scheduler/main.go
```

### API Endpoints

The server provides the following endpoints:

- `GET /api/v1/health` - Health check
- `GET /api/v1/ready` - Readiness probe
- `GET /api/v1/live` - Liveness probe
- `POST /api/v1/auth/login` - User authentication
- `POST /api/v1/auth/refresh` - Token refresh
- `GET /api/v1/users` - List users
- `POST /api/v1/users` - Create user
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Adding New Features

1. **Define Domain Model** (`internal/domain/entity/`)
```go
type Product struct {
    ID          uuid.UUID
    Name        string
    Price       decimal.Decimal
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

2. **Create Repository Interface** (`internal/domain/repository/`)
```go
type ProductRepository interface {
    Save(ctx context.Context, product *entity.Product) error
    FindByID(ctx context.Context, id uuid.UUID) (*entity.Product, error)
    FindAll(ctx context.Context) ([]*entity.Product, error)
}
```

3. **Implement Repository** (`internal/infrastructure/repository/`)
```go
type productRepoImpl struct {
    db *sqlx.DB
}

func (r *productRepoImpl) Save(ctx context.Context, product *entity.Product) error {
    // Implementation
}
```

4. **Create Application Service** (`internal/application/command_service/`)
```go
type ProductCommandService struct {
    productRepo repository.ProductRepository
    logger      *zap.Logger
}

func (s *ProductCommandService) CreateProduct(ctx context.Context, cmd CreateProductCommand) error {
    // Business logic
}
```

5. **Add HTTP Controller** (`internal/interfaces/http/controller/`)
```go
type ProductController struct {
    productCmdService *command_service.ProductCommandService
}

func (c *ProductController) CreateProduct(ctx context.Context, req *app.RequestContext) {
    // Handle HTTP request
}
```

6. **Register Routes** (`internal/interfaces/http/router.go`)
```go
productController := container.ProductController()
base.POST("/products", productController.CreateProduct)
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./domain/...
```

### Building

```bash
# Build all binaries
./scripts/build.sh

# Build specific binary
go build -o bin/server ./cmd/server
go build -o bin/migrate ./cmd/migrate
go build -o bin/scheduler ./cmd/scheduler
```

## Docker Support

### Build and Run

```bash
# Build image
docker build -t go-template .

# Run container
docker run -p 8080:8080 \
  -e APP_ENV=prod \
  -v $(pwd)/config.prod.yaml:/app/config.prod.yaml \
  go-template
```

### Docker Compose

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

## Deployment

### Environment Variables

- `APP_ENV`: Environment (dev, test, prod)
- `HTTP_PORT`: HTTP server port
- `DB_DSN`: Database connection string
- `REDIS_URL`: Redis connection URL
- `JWT_SECRET`: JWT signing secret

### Health Checks

- **Health**: `GET /api/v1/health` - Overall application health
- **Readiness**: `GET /api/v1/ready` - Ready to serve traffic
- **Liveness**: `GET /api/v1/live` - Application is alive

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Code Style

- Follow Go conventions and best practices
- Use `gofmt` for code formatting
- Write comprehensive tests
- Document public APIs
- Follow Clean Architecture principles

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Maintainer

- [@lyonnee](https://github.com/lyonnee)

## Acknowledgments

- [CloudWeGo](https://github.com/cloudwego) for the excellent Hertz framework
- [Uber](https://github.com/uber-go) for the Zap logging library
- Clean Architecture community for architectural guidance
- All contributors who helped shape this project