<div align="center">
</br>

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
go run cmd/server/main.go -e dev
```

### Environment Variables

Set the runtime environment through:
- Command line flags: `-e` or `--env`
- Environment variable: `APP_ENV`
- Default value: `prod`

## Features

- 📦 Clean Architecture
- 🎯 Domain-Driven Design (DDD)
- 🔐 JWT Authentication
- 📝 Structured Logging (Zap)
- 🗄️ Any Database
- 🔄 Graceful Shutdown
- 🐳 Docker Support
- 📊 Prometheus Metrics
- ⚡ High-Performance HTTP Router (Fiber)
- 🔌 gRPC Support
- 📨 Message Queue Integration
- 💾 Caching Support
- ⏰ Task Scheduling
- 🔧 Multi-Environment Configuration

## Project Structure

```
go-template/                          # Project root
├── cmd/                             # Command line entry points
│   ├── migrate/                     # Database migration tool
│   ├── scheduler/                   # Task scheduler
│   └── server/                      # Main server
│
├── sql/                             # SQL definition files
│
├── config/                          # Configuration management
│   ├── config.go                    # Config loading logic
│   └── model.go                     # Config structure definitions
│
├── internal/                        # Internal application code
│   ├── application/                 # Application layer
│   │   ├── auth_service.go         # Auth service interface
│   │   └── impl/                   # Service implementations
│   ├── domain/                     # Domain layer
│   ├── infrastructure/            # Infrastructure layer
│   │   └── repositories/         # Repository implementations
│   └── interfaces/               # Interface layer
│       ├── grpc/                # gRPC interfaces
│       └── http/                # HTTP interfaces
│           ├── controller/      # Controllers
│           ├── dto/            # Data Transfer Objects
│           └── middleware/     # Middlewares
│
├── pkg/                        # Public packages
│   ├── modules/               # Core modules
│   │   ├── auth/            # Authentication
│   │   ├── blockchain/     # Blockchain integration
│   │   ├── cache/         # Caching
│   │   ├── logger/        # Logging
│   │   ├── mq/           # Message Queue
│   │   └── persistence/  # Data persistence
│   └── scheduler/         # Scheduler module
│
└── servers/               # Server implementations
    ├── http.go           # HTTP server
    └── rpc.go            # RPC server
```

## Architecture

The project follows Clean Architecture and DDD principles, organized in layers:

1. **Interface Layer**
   - HTTP/gRPC request handling
   - Request validation
   - Response formatting

2. **Application Layer**
   - Business process orchestration
   - Domain object coordination
   - Transaction management

3. **Domain Layer**
   - Core business logic
   - Domain models
   - Domain services

4. **Infrastructure Layer**
   - Data persistence
   - Message queuing
   - Caching
   - Logging
   - Authentication

## Development Guide

### Adding New Features

1. Define domain models in `internal/domain`
2. Implement business logic in `internal/application`
3. Add API endpoints in `internal/interfaces`
4. Implement infrastructure in `internal/infrastructure`

### Configuration Management

Configuration files in project root:
- `config.dev.yaml`: Development environment
- `config.test.yaml`: Testing environment
- `config.prod.yaml`: Production environment

### Database Migration

```bash
go run cmd/migrate/main.go
```

### Running Tests

```bash
go test ./...
```

## Docker Support

Build image:
```bash
docker build -t your-app-name .
```

Run container:
```bash
docker run -p 8080:8080 your-app-name
```

## Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Maintainer

- [@lyonnee](https://github.com/lyonnee)

## Acknowledgments

Thanks to all contributors who helped shape this project.