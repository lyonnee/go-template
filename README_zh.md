# Go 项目模板

一个基于 Go 语言的现代化 Web 服务项目模板，采用领域驱动设计(DDD)和整洁架构(Clean Architecture)思想构建。

| [English](README.md) | 中文 |
| --- | --- |

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

- 📦 **整洁架构** - 清晰的分层架构设计，关注点分离
- 🎯 **领域驱动设计** - 丰富的领域模型和业务逻辑封装
- 🔐 **认证系统** - JWT 和 OAuth 集成就绪
- 📝 **日志基础设施** - 结构化日志，支持多种输出
- 🗄️ **数据库支持** - 多数据库兼容性
- 💾 **缓存层** - Redis 集成
- 🔄 **优雅关机** - 合适的资源清理
- 🐳 **Docker 就绪** - 包含多阶段构建
- ⚡ **高性能 HTTP** - CloudWeGo Hertz 框架
- 🔌 **gRPC 支持** - Protocol Buffer 集成
- 📨 **消息队列** - 异步处理基础设施
- ⏰ **任务调度** - 定时任务支持
- 🔧 **多环境配置** - 开发、测试、生产配置
- 🏗️ **依赖注入** - IoC 容器包含
- 🆔 **ID 生成** - 分布式 ID 生成
- 🧪 **测试结构** - 测试组织和工具
- 📨 **消息队列集成** - 异步消息处理
- ⏰ **定时任务调度** - 支持 Cron 表达式
- 🔧 **多环境配置** - 基于 YAML 的配置管理
- 🏗️ **依赖注入** - Samber/do IoC 容器
- 🆔 **ID 生成** - 基于 Snowflake 的唯一 ID 生成
- � **密码哈希** - bcrypt 安全密码存储
- 🔗 **区块链集成** - 内置区块链工具
- 🧪 **测试支持** - 完整的测试工具和结构

## 项目结构

本模板遵循整洁架构和 DDD 原则，具有以下结构：

```
go-template/
├── Dockerfile                        # 容器定义
├── LICENSE                           # 许可证
├── go.mod                            # Go 模块定义
├── go.sum                            # 依赖校验和
├── main.go                           # 程序入口
├── README.md                         # 英文说明
├── README_zh.md                      # 中文说明
├── _logs/                            # 本地日志输出
│   └── dev.log
│
├── application/
│   └── cron/                         # 外层调度器入口/封装
│       └── scheduler.go
│
├── configs/                          # 多环境配置文件
│   ├── config.dev.yaml
│   ├── config.prod.yaml
│   └── config.test.yaml
│
├── infrastructure/
│   └── di/                           # 根级依赖注入装配
│       └── injector.go
│
├── internal/                         # 业务实现（遵循 Go internal 隔离）
│   ├── application/                  # 应用层（用例编排）
│   │   ├── commands/                 # 写模型（命令）
│   │   │   ├── auth_command_service.go
│   │   │   └── user_command_service.go
│   │   ├── queries/                  # 读模型（查询）
│   │   │   └── user_query_service.go
│   │   └── scheduler/                # 定时任务编排
│   │       ├── scheduler.go
│   │       └── jobs/
│   │           └── test_job.go
│   │
│   ├── domain/                       # 领域层
│   │   ├── entity/                   # 领域实体
│   │   │   └── user.go
│   │   ├── errors/                   # 领域错误
│   │   │   └── user_errors.go
│   │   ├── repository/               # 仓储接口
│   │   │   ├── eth_repository.go
│   │   │   └── user_repository.go
│   │   └── service/                  # 领域服务
│   │       ├── infra_service.go
│   │       └── user_service.go
│   │
│   ├── infrastructure/               # 基础设施实现
│   │   ├── auth/                     # 认证/JWT/OAuth
│   │   │   ├── auth.go
│   │   │   ├── jwt.go
│   │   │   └── oauth.go
│   │   ├── blockchain/               # 区块链相关
│   │   │   └── blockchain.go
│   │   ├── cache/                    # 缓存与 Redis
│   │   │   ├── cache.go
│   │   │   ├── keys.go
│   │   │   └── redis.go
│   │   ├── config/                   # 配置装载
│   │   │   ├── config.go
│   │   │   └── types.go
│   │   ├── database/                 # 数据库访问
│   │   │   ├── database.go
│   │   │   ├── executor.go
│   │   │   ├── logger.go
│   │   │   └── postgres.go
│   │   ├── mq/                       # 消息队列
│   │   │   └── mq.go
│   │   └── repository_impl/              # 持久化实现
│   │       ├── user_repository.go
│   │       └── model/
│   │           ├── base_model.go
│   │           └── user.go
│   │
│   └── interfaces/                   # 适配层（对外接口）
│       ├── event_handler/            # 事件处理
│       │   └── event_handler.go
│       ├── grpc/                     # gRPC 定义
│       │   └── user.proto
│       └── http/                     # HTTP 接口
│           ├── controller/           # 控制器
│           │   ├── auth_controller.go
│           │   ├── health_controller.go
│           │   └── user_controller.go
│           ├── dto/                  # DTO 定义
│           │   ├── auth.go
│           │   ├── base_response.go
│           │   ├── pagequery.go
│           │   └── user.go
│           ├── middleware/           # 中间件
│           │   ├── cors.go
│           │   ├── jwt.go
│           │   ├── logger.go
│           │   ├── recovery.go
│           │   └── trace.go
│           └── router.go             # 路由
│
├── pkg/                              # 通用库
│   ├── di/                           # DI 帮助
│   │   └── injector.go
│   ├── idgen/                        # ID 生成
│   │   └── id_generator.go
│   ├── log/                          # 日志封装
│   │   ├── log.go
│   │   └── zap_logger.go
│   └── util/                         # 工具方法
│       └── bcrypt.go
│
├── scripts/                          # 构建与启动脚本
│   ├── build.sh
│   └── start.sh
│
├── services/                         # 服务启动入口（HTTP/gRPC/Cron）
│   ├── cron.go
│   ├── grpc.go
│   ├── http.go
│   └── service.go
│
├── sqls/                             # 数据库初始化/迁移 SQL
│   └── user.sql
│
└── test/                             # 测试目录
```

## 架构说明

项目采用整洁架构（Clean Architecture）和领域驱动设计（DDD）思想构建，分为以下几层：

### 1. 接口层 (Interfaces Layer)
- **HTTP 控制器**: 处理 HTTP 请求和响应
- **gRPC 服务**: 处理 RPC 调用
- **中间件**: 请求拦截和处理（认证、日志、CORS 等）
- **DTO**: 数据传输对象，用于接口层数据交换

### 2. 应用层 (Application Layer)
- **命令执行器**: 处理写操作（CQRS 模式）
- **查询执行器**: 处理读操作（CQRS 模式）
- **应用服务**: 编排领域对象，处理业务流程
- **事务管理**: 确保数据一致性

### 3. 领域层 (Domain Layer)
- **实体**: 具有唯一标识的业务对象
- **值对象**: 不可变的业务概念
- **领域服务**: 跨实体的业务逻辑
- **仓储接口**: 数据访问抽象
- **领域事件**: 业务事件定义

### 4. **基础设施层**（技术细节）
- **仓储实现**: 数据持久化实现
- **缓存实现**: 缓存策略
- **消息队列**: 异步通信
- **配置**: 环境特定设置

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

func (c *ProductController) CreateProduct(ctx context.Context, req *app.RequestContext) {
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

func (c *ProductController) CreateProduct(ctx context.Context, req *app.RequestContext) {
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

本模板使用统一的依赖注入接口，基于 samber/do/v2 实现。**所有服务（仓储、领域服务、应用服务、控制器）都必须通过 `injector.go` 提供的接口进行注册。**

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
func RateLimit() app.HandlerFunc {
    return func(ctx context.Context, c *app.RequestContext) {
        // 限流逻辑
        c.Next(ctx)
    }
}
```

#### 2. 注册中间件
更新路由器以使用中间件：

```go
// internal/interfaces/http/router.go
h.Use(middleware.RateLimit())
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

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 遵循模板结构和架构原则
4. 为你的更改编写测试
5. 提交改动 (`git commit -m 'Add amazing feature'`)
6. 推送分支 (`git push origin feature/amazing-feature`)
7. 创建 Pull Request

## 许可证

本项目使用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 维护者

- [@lyonnee](https://github.com/lyonnee)