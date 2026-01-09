<div align="center">

# Go 项目模板

| [English](README.md) | 中文 |
| --- | --- |

一个基于 Go 语言的现代化 Web 服务项目模板，采用领域驱动设计(DDD)和整洁架构(Clean Architecture)原则构建。
</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/lyonnee/go-template)](https://goreportcard.com/report/github.com/lyonnee/go-template)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lyonnee/go-template)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## 快速开始

### 使用模板创建项目

1. 安装 gonew 工具
```bash
go install golang.org/x/tools/cmd/gonew@latest
```

2. 使用模板创建新项目
```bash
gonew github.com/lyonnee/go-template github.com/your-username/your-project
```

3. 进入项目目录
```bash
cd your-project
```

4. 启动开发服务器
```bash
go run . -env dev
```

## 模板特性

本模板提供了生产就绪的 Go Web 服务，包含：

### 🏗️ 架构设计
- 📦 **整洁架构** - 严格遵循 Clean Architecture 原则，清晰的分层设计
- 🎯 **领域驱动设计** - DDD 实践，丰富的领域模型和业务逻辑封装
- 🔌 **CQRS 模式** - 命令查询职责分离，读写模型分离
- 🏗️ **依赖注入** - 基于 samber/do/v2 的 IoC 容器

### 🚀 技术栈
- ⚡ **高性能 HTTP** - Gin Web 框架
- 🗄️ **数据库支持** - PostgreSQL + SQLx，支持事务管理
- 💾 **缓存层** - Redis 集成，支持缓存上下文
- 🔐 **认证系统** - JWT + OAuth 完整实现
- 📝 **结构化日志** - Zap + Lumberjack 日志系统
- ⏰ **任务调度** - Cron 定时任务支持
- 🆔 **ID 生成** - Snowflake 分布式 ID 生成

### 🔧 开发支持
- 🔧 **多环境配置** - dev/test/prod 环境配置管理
- 🐳 **Docker 就绪** - 多阶段构建 Dockerfile
- 🔄 **优雅关机** - 完整的资源清理机制
- 📊 **健康检查** - 完整的健康检查端点
- 🛡️ **中间件系统** - CORS、Recovery、JWT 等中间件
- 🔒 **密码安全** - bcrypt 密码哈希存储

## 项目结构

本模板严格遵循整洁架构 (Clean Architecture) 和领域驱动设计 (DDD) 原则：

```
go-template/
├── main.go                           # 程序入口点
├── go.mod                            # Go 模块定义
├── go.sum                            # 依赖校验和
├── Dockerfile                        # 容器化部署
├── _logs/                            # 本地日志目录
│   └── dev.log
│
├── configs/                          # 多环境配置
│   ├── config.dev.yaml              # 开发环境配置
│   ├── config.test.yaml              # 测试环境配置
│   └── config.prod.yaml              # 生产环境配置
│
├── internal/                         # 核心业务代码 (Go internal 包)
│   ├── application/                  # 应用层 - 用例编排
│   │   ├── commands/                 # CQRS 写侧 (命令)
│   │   │   ├── auth_command_service.go
│   │   │   └── user_command_service.go
│   │   ├── queries/                  # CQRS 读侧 (查询)
│   │   │   └── user_query_service.go
│   │   └── scheduler/                # 任务调度
│   │       ├── scheduler.go
│   │       └── jobs/
│   │           └── test_job.go
│   │
│   ├── domain/                       # 领域层 - 业务逻辑核心
│   │   ├── entity/                   # 领域实体
│   │   │   └── user.go
│   │   ├── errors/                   # 领域错误定义
│   │   │   └── user_errors.go
│   │   ├── repository/               # 仓储接口 (端口)
│   │   │   ├── repository.go
│   │   │   ├── user_repository.go
│   │   │   └── eth_repository.go
│   │   └── service/                  # 领域服务
│   │       ├── user_service.go
│   │       └── infra_service.go
│   │
│   ├── infrastructure/               # 基础设施层 - 技术实现
│   │   ├── config/                   # 配置管理
│   │   │   ├── config.go
│   │   │   └── types.go
│   │   ├── database/                 # 数据库基础设施
│   │   │   ├── database.go
│   │   │   ├── dbcontext.go
│   │   │   ├── logger.go
│   │   │   └── postgres.go
│   │   ├── cache/                    # 缓存基础设施
│   │   │   ├── cache.go
│   │   │   ├── cachecontext.go
│   │   │   ├── keys.go
│   │   │   └── redis.go
│   │   ├── auth/                     # 认证基础设施
│   │   │   ├── auth.go
│   │   │   ├── jwt.go
│   │   │   └── oauth.go
│   │   ├── repository_impl/          # 仓储实现 (适配器)
│   │   │   ├── user_repository.go
│   │   │   └── model/                # 数据库模型
│   │   │       ├── base_model.go
│   │   │       └── user.go
│   │   ├── blockchain/               # 区块链工具
│   │   │   └── blockchain.go
│   │   └── mq/                       # 消息队列
│   │       └── mq.go
│   │
│   └── interfaces/                   # 接口适配层 - 外部接口
│       ├── http/                     # HTTP 接口
│       │   ├── router.go             # 路由配置
│       │   ├── controller/           # HTTP 控制器
│       │   │   ├── auth_controller.go
│       │   │   ├── user_controller.go
│       │   │   └── health_controller.go
│       │   ├── dto/                  # 数据传输对象
│       │   │   ├── base_response.go
│       │   │   ├── auth.go
│       │   │   ├── user.go
│       │   │   └── pagequery.go
│       │   └── middleware/           # HTTP 中间件
│       │       ├── cors.go
│       │       ├── jwt.go
│       │       ├── logger.go
│       │       ├── recovery.go
│       │       └── trace.go
│       ├── grpc/                     # gRPC 接口
│       └── event_handler/            # 事件处理器
│           └── event_handler.go
│
├── pkg/                              # 共享工具库
│   ├── di/                           # 依赖注入容器
│   │   └── injector.go
│   ├── log/                          # 日志工具
│   │   ├── log.go
│   │   └── zap_logger.go
│   ├── idgen/                        # ID 生成器
│   │   └── id_generator.go
│   └── util/                         # 通用工具
│       └── bcrypt.go
│
├── services/                         # 服务启动管理
│   ├── services.go                   # 服务注册器
│   ├── http.go                       # HTTP 服务
│   ├── grpc.go                       # gRPC 服务
│   └── cron.go                       # 定时任务服务
│
├── scripts/                          # 构建脚本
│   ├── build.sh
│   └── start.sh
│
├── sqls/                             # 数据库迁移
│   └── user.sql
│
└── test/                             # 测试目录
```

## 架构说明

项目采用整洁架构（Clean Architecture）和领域驱动设计（DDD）思想构建，分为以下几层：

### 架构分层说明

#### 1. 接口适配层 (Interfaces Layer)
- **HTTP 控制器**: 处理 RESTful API 请求响应，负责参数绑定和响应格式化
- **中间件系统**: 横切关注点实现 (认证、日志、CORS、恢复等)
- **DTO 对象**: 数据传输对象，用于接口层与应用层的数据交换
- **路由管理**: 统一的路由注册和管理

#### 2. 应用层 (Application Layer)
- **命令服务**: CQRS 写侧，处理业务操作和状态变更
- **查询服务**: CQRS 读侧，处理数据查询和展示逻辑
- **任务调度**: 定时任务和后台作业的编排管理
- **事务协调**: 跨领域对象的事务管理和数据一致性

#### 3. 领域层 (Domain Layer)
- **领域实体**: 具有唯一标识的核心业务对象，包含业务逻辑
- **领域服务**: 不适合放在单个实体中的业务逻辑
- **仓储接口**: 数据访问的抽象定义 (端口)
- **领域错误**: 业务相关的错误定义和处理

#### 4. 基础设施层 (Infrastructure Layer)
- **仓储实现**: 具体的数据持久化实现 (适配器)
- **数据库访问**: PostgreSQL 连接、事务管理、查询日志
- **缓存实现**: Redis 缓存策略和上下文管理
- **认证系统**: JWT Token 生成验证、OAuth 集成
- **配置管理**: 多环境配置加载和类型安全访问

## 开发指南

### 添加新的业务功能

#### 1. 定义领域实体
在 `internal/domain/entity/` 中创建新的业务实体：

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

#### 2. 创建仓储接口
在 `internal/domain/repository/` 中定义数据访问接口：

```go
// domain/repository/product_repository.go
type ProductRepository interface {
    Save(ctx context.Context, product *entity.Product) error
    FindByID(ctx context.Context, id int64) (*entity.Product, error)
    FindAll(ctx context.Context) ([]*entity.Product, error)
    Delete(ctx context.Context, id int64) error
}
```

#### 3. 实现仓储
在 `internal/infrastructure/repository_impl/` 中创建具体实现：

```go
// internal/infrastructure/repository_impl/product_repository.go
type ProductRepoImpl struct {
    db *sqlx.DB
}

func (r *ProductRepoImpl) Save(ctx context.Context, product *entity.Product) error {
    // 数据库实现
}
```

#### 4. 注册仓储到依赖容器
在同一个文件中使用 `init()` 函数注册仓储：

```go
// internal/infrastructure/repository_impl/product_repository.go
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
    // 数据库实现
}
```

#### 4. 创建应用服务
在 `internal/application/commands/` 中实现写模型业务逻辑（示例）：

```go
// internal/application/commands/product_command_service.go
type ProductService struct {
    productRepo repository.ProductRepository
}

func (s *ProductService) CreateProduct(ctx context.Context, req CreateProductRequest) error {
    // 业务逻辑实现
}
```

#### 5. 注册应用服务到依赖容器
在同一个文件中注册应用服务：

```go
// internal/application/commands/product_command_service.go
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
    // 业务逻辑实现
}
```

#### 5. 添加 HTTP 控制器
在 `internal/interfaces/http/controller/` 中处理 HTTP 请求：

```go
// internal/interfaces/http/controller/product_controller.go
type ProductController struct {
    productService *service.ProductService
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
    // HTTP 请求处理
}
```

#### 6. 注册控制器到依赖容器
在同一个文件中注册控制器：

```go
// internal/interfaces/http/controller/product_controller.go
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

func (c *ProductController) CreateProduct(ctx *gin.Context) {
    // HTTP 请求处理
}
```

#### 6. 注册路由
在 `internal/interfaces/http/router.go` 中更新路由：

```go
// 添加到 router.go
productController := di.Get[*ProductController]()
v1.POST("/products", productController.CreateProduct)
v1.GET("/products/:id", productController.GetProduct)
v1.PUT("/products/:id", productController.UpdateProduct)
v1.DELETE("/products/:id", productController.DeleteProduct)
```

### 添加新的配置项

#### 1. 更新配置结构
在 `internal/infrastructure/config/types.go` 中添加新的配置部分：

```go
type Config struct {
    // ... 现有字段
    Email EmailConfig `yaml:"email"`
}

type EmailConfig struct {
    Provider  string `yaml:"provider"`
    SMTPHost  string `yaml:"smtp_host"`
    SMTPPort  int    `yaml:"smtp_port"`
    Username  string `yaml:"username"`
    Password  string `yaml:"password"`
}
```

#### 2. 更新配置文件
在环境配置文件中添加配置：

```yaml
# configs/config.dev.yaml
email:
  provider: smtp
  smtp_host: smtp.gmail.com
  smtp_port: 587
  username: your-email@gmail.com
  password: your-password
```

#### 3. 注册服务
使用统一的依赖注入接口注册：

```go
// 在适当的初始化位置注册（例如 internal/infrastructure/email/email_service.go）
di.AddSingleton(func() (EmailService, error) {
    config := di.Get[*config.Config]()
    return &emailServiceImpl{
        config: config.Email,
    }, nil
})
```

### 依赖注入使用指南

本项目采用自研的依赖注入系统，基于 samber/do/v2 封装。**所有服务 (仓储、领域服务、应用服务、控制器) 都必须通过统一的 DI 接口进行注册。**

#### 核心原则

1. **自注册模式**：服务使用 `init()` 函数自己注册到容器
2. **工厂模式**：使用 `New` 函数作为服务工厂
3. **类型安全**：利用 Go 泛型确保类型安全
4. **统一接口**：所有依赖注册都使用 `di.AddSingleton` 或 `di.AddTransient`

#### 服务注册模式

每个服务都应该遵循这个模式：

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
    // 业务逻辑实现
}
```

#### 仓储注册

```go
// internal/infrastructure/repository_impl/user_repository.go
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

#### 应用服务注册

```go
// internal/application/commands/user_command_service.go
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

#### 控制器注册

```go
// internal/interfaces/http/controller/user_controller.go
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

#### 瞬态服务注册

对于需要每次都创建新实例的服务：

```go
// internal/infrastructure/email/email_service.go
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

#### 获取服务

```go
// 在路由或其他组件中获取服务
func SetupRoutes() {
    userController := di.Get[*controller.UserController]()
    v1.POST("/users", userController.CreateUser)
}

// 在服务中获取依赖
func (s *SomeService) ProcessUser() {
    userRepo := di.Get[repository.UserRepository]()
    // 使用仓储...
}
```

#### 重要最佳实践

- **自注册**：每个服务在自己的 `init()` 函数中注册
- **工厂函数**：始终提供 `New` 函数作为服务工厂
- **接口注册**：仓储和领域服务优先注册接口类型
- **具体注册**：应用服务和控制器使用具体类型
- **依赖注入**：在工厂函数中始终使用 `di.Get[T]()` 解析依赖
- **错误处理**：工厂函数应该返回 `(T, error)` 以便正确处理错误

### 添加新的中间件

#### 1. 创建中间件
在 `internal/interfaces/http/middleware/` 中添加新中间件：

```go
// internal/interfaces/http/middleware/rate_limit.go
func RateLimit() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 限流逻辑
        c.Next()
    }
}
```

#### 2. 注册中间件
更新路由器以使用中间件：

```go
// internal/interfaces/http/router.go
engine.Use(middleware.RateLimit())
```

### 添加新的服务

#### 1. 创建服务接口
在 `internal/domain/service/` 中定义服务契约：

```go
// domain/service/notification_service.go
type NotificationService interface {
    SendEmail(ctx context.Context, to, subject, body string) error
    SendSMS(ctx context.Context, to, message string) error
}
```

#### 2. 实现领域服务
在 `internal/domain/service/` 中创建实现：

```go
// internal/domain/service/notification_service_impl.go
type NotificationServiceImpl struct {
    config EmailConfig
}

func (s *NotificationServiceImpl) SendEmail(ctx context.Context, to, subject, body string) error {
    // 邮件发送实现
}
```

#### 3. 注册领域服务到依赖容器
在同一个文件中实现和注册领域服务：

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
    // 邮件发送实现
}
```

### 添加数据库模型

#### 1. 创建数据库模型
在 `internal/infrastructure/repository_impl/model/` 中添加模型：

```go
// internal/infrastructure/repository_impl/model/product.go
type Product struct {
    BaseModel
    Name  string  `db:"name"`
    Price float64 `db:"price"`
}
```

#### 2. 创建迁移
在 `sqls/` 中添加 SQL 文件：

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

### 测试你的更改

```bash
# 运行所有测试
go test ./...

# 运行特定包测试
go test ./internal/domain/...
go test ./internal/application/...

# 运行覆盖率测试
go test -cover ./...
```

### 构建和运行

```bash
# 构建应用程序
./scripts/build.sh

# 运行开发服务器
go run . -env dev

# 使用 Docker 运行
docker build -t your-app .
docker run -p 8080:8080 your-app
```

## API 接口说明

### 健康检查接口
```
GET /api/health      # 健康状态检查
GET /api/ready       # 就绪状态检查  
GET /api/live        # 存活状态检查
```

### 认证接口
```
POST /api/auth/login    # 用户登录
POST /api/auth/refresh  # 刷新令牌
```

### 用户接口
```
POST /api/users             # 用户注册
GET  /api/users/:id         # 获取用户信息 (需认证)
PUT  /api/users/:id/username # 更新用户名 (需认证)
```

## 环境要求

- **Go**: 1.23.7 或更高版本
- **PostgreSQL**: 12 或更高版本  
- **Redis**: 6.0 或更高版本
- **Docker**: 20.10 或更高版本 (可选)

## 配置说明

项目支持多环境配置，通过 `-env` 参数指定：

```bash
# 开发环境 (默认)
go run . -env dev

# 测试环境
go run . -env test  

# 生产环境
go run . -env prod
```

配置文件位于 `configs/` 目录：
- `config.dev.yaml` - 开发环境配置
- `config.test.yaml` - 测试环境配置
- `config.prod.yaml` - 生产环境配置

## 贡献指南

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 遵循项目的架构原则和代码规范
4. 为新功能编写测试用例
5. 确保所有测试通过
6. 提交更改 (`git commit -m 'Add amazing feature'`)
7. 推送分支 (`git push origin feature/amazing-feature`)
8. 创建 Pull Request

### 代码规范
- 遵循 Go 官方代码风格
- 使用有意义的变量和函数名
- 为公共接口提供文档注释
- 保持函数简洁，单一职责

## 许可证

本项目使用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 维护者

- [@lyonnee](https://github.com/lyonnee)