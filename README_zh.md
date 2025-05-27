# Go 项目模板

一个基于 Go 语言的现代化 Web 服务项目模板，采用领域驱动设计(DDD)和整洁架构(Clean Architecture)思想构建。

| [English](README.md) | 中文 |
| --- | --- |

## 快速开始

### 使用 gonew 创建项目

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
go run cmd/server/main.go -e dev
```

### 环境变量

你可以通过以下方式设置运行环境：
- 命令行参数: `-e` 或 `--env`
- 环境变量: `APP_ENV`
- 默认值: `prod`

## 项目特性

- 📦 **整洁架构** (Clean Architecture) - 清晰的分层架构设计
- 🎯 **领域驱动设计** (DDD) - 以业务领域为核心的设计方法
- 🔐 **JWT 认证** - 支持访问令牌和刷新令牌
- 📝 **结构化日志** (Zap) - 高性能的结构化日志记录
- 🗄️ **多数据库支持** - PostgreSQL 和 MySQL
- 🔄 **优雅关机** - 信号处理和资源清理
- 🐳 **Docker 支持** - 多阶段构建和容器化部署
- ⚡ **高性能 HTTP 框架** (CloudWeGo Hertz) - 字节跳动开源的高性能 HTTP 框架
- 🔌 **gRPC 支持** - 高性能 RPC 通信
- 📨 **消息队列集成** - 异步消息处理
- 💾 **Redis 缓存** - 高性能缓存支持
- ⏰ **定时任务调度** - 支持 Cron 表达式
- 🔧 **多环境配置** - 开发、测试、生产环境配置
- 🏗️ **依赖注入** - 基于容器的依赖管理
- 📊 **CQRS 模式** - 命令查询职责分离
- 🔍 **中间件支持** - 恢复、CORS、追踪、日志等
- 🛡️ **密码加密** - bcrypt 密码哈希
- 🔗 **区块链集成** - 区块链相关功能支持

## 项目结构

```
go-template/                          # 项目根目录
├── cmd/                             # 命令行入口目录
│   ├── migrate/                     # 数据库迁移工具
│   │   └── main.go                  # 迁移命令入口
│   ├── schduler/                    # 定时任务调度器
│   │   └── main.go                  # 调度器入口
│   └── server/                      # 主服务器
│       └── main.go                  # 服务器入口
│
├── config/                          # 配置管理模块
│   ├── auth.go                      # 认证配置
│   ├── cache.go                     # 缓存配置
│   ├── config.go                    # 配置加载逻辑
│   ├── http.go                      # HTTP 服务配置
│   ├── log.go                       # 日志配置
│   └── persistence.go               # 持久化配置
│
├── internal/                        # 内部应用代码
│   ├── app/                         # 应用容器
│   │   └── container.go             # 依赖注入容器
│   ├── application/                 # 应用层：处理业务流程
│   │   ├── command_executor/        # 命令执行器 (CQRS)
│   │   └── query_executor/          # 查询执行器 (CQRS)
│   ├── domain/                      # 领域层：核心业务逻辑
│   │   ├── entity/                  # 领域实体
│   │   ├── errors/                  # 领域错误
│   │   ├── repository/              # 仓储接口
│   │   ├── user_domain_service.go   # 用户领域服务
│   │   └── valueobject/             # 值对象
│   ├── infrastructure/              # 基础设施适配层
│   │   ├── cache/                   # 缓存实现
│   │   ├── eventbus/                # 事件总线
│   │   ├── log/                     # 日志实现
│   │   ├── repository/              # 仓储实现
│   │   └── scheduler/               # 调度器实现
│   └── interfaces/                  # 接口层
│       ├── grpc/                    # gRPC 接口
│       └── http/                    # HTTP 接口
│           ├── controller/          # 控制器
│           ├── dto/                 # 数据传输对象
│           ├── middleware/          # 中间件
│           └── router.go            # 路由配置
│
├── pkg/                             # 公共包
│   ├── auth/                        # 认证模块
│   │   ├── jwt.go                   # JWT 实现
│   │   ├── oauth.go                 # OAuth 实现
│   │   └── password.go              # 密码处理
│   ├── blockchain/                  # 区块链集成
│   │   └── blockchain.go            # 区块链功能
│   ├── cache/                       # 缓存模块
│   │   └── cache.go                 # 缓存接口
│   ├── hash/                        # 哈希工具
│   ├── log/                         # 日志模块
│   │   ├── zap_logger.go            # Zap 日志器
│   │   └── zap_sugar_logger.go      # Zap Sugar 日志器
│   ├── mq/                          # 消息队列
│   │   └── mq.go                    # 消息队列接口
│   └── persistence/                 # 数据持久化
│       ├── persistence.go           # 持久化接口
│       └── postgres.go              # PostgreSQL 实现
│
├── server/                          # 服务器实现
│   ├── http.go                      # HTTP 服务器
│   └── rpc.go                       # RPC 服务器
│
├── scripts/                         # 脚本目录
│   ├── build.sh                     # 构建脚本
│   └── start.sh                     # 启动脚本
│
├── sql/                             # SQL 文件
│   └── user.sql                     # 用户表结构
│
├── test/                            # 测试目录
│
├── config.dev.yaml                  # 开发环境配置
├── config.test.yaml                 # 测试环境配置
├── config.prod.yaml                 # 生产环境配置
├── Dockerfile                       # Docker 构建文件
└── go.mod                           # Go 模块文件
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

### 4. 基础设施层 (Infrastructure Layer)
- **仓储实现**: 数据持久化具体实现
- **缓存实现**: Redis 缓存服务
- **日志实现**: 结构化日志记录
- **事件总线**: 事件发布和订阅
- **调度器**: 定时任务执行

## 技术栈

- **Web 框架**: [CloudWeGo Hertz](https://github.com/cloudwego/hertz) - 高性能 HTTP 框架
- **配置管理**: [Viper](https://github.com/spf13/viper) - 配置文件解析
- **日志**: [Zap](https://github.com/uber-go/zap) - 高性能结构化日志
- **数据库**: PostgreSQL/MySQL
- **缓存**: Redis
- **认证**: JWT (JSON Web Tokens)
- **密码加密**: bcrypt
- **测试**: Go 标准测试库 + [Testify](https://github.com/stretchr/testify)
- **Mock**: [GoMock](https://github.com/golang/mock)
- **容器化**: Docker

## 配置管理

项目使用 Viper 进行配置管理，支持多环境配置：

### 配置文件结构

```yaml
# config.dev.yaml 示例
http:
  port: 8081

log:
  level: debug
  format: json
  output: stdout
  file:
    enabled: true
    path: ./_logs/app.log
    max_size: 100
    max_backups: 3
    max_age: 28
    compress: true

auth:
  jwt:
    secret: your-secret-key
    access_token_expire: 15m
    refresh_token_expire: 168h

persistence:
  database:
    driver: postgres
    host: localhost
    port: 5432
    username: postgres
    password: password
    database: go_template
    ssl_mode: disable
    max_open_conns: 25
    max_idle_conns: 5
    conn_max_lifetime: 300s

cache:
  redis:
    host: localhost
    port: 6379
    password: ""
    db: 0
    pool_size: 10
    min_idle_conns: 5
```

### 环境配置

- `config.dev.yaml`: 开发环境配置
- `config.test.yaml`: 测试环境配置  
- `config.prod.yaml`: 生产环境配置

## 开发指南

### 添加新功能

1. **定义领域模型** (`internal/domain/entity/`)
```go
type User struct {
    ID       string
    Username string
    Email    string
    // ... 其他字段
}
```

2. **定义仓储接口** (`internal/domain/repository/`)
```go
type UserRepository interface {
    Save(ctx context.Context, user *entity.User) error
    FindByID(ctx context.Context, id string) (*entity.User, error)
}
```

3. **实现应用服务** (`internal/application/`)
```go
type UserService struct {
    userRepo repository.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, cmd CreateUserCommand) error {
    // 业务逻辑实现
}
```

4. **添加 HTTP 控制器** (`internal/interfaces/http/controller/`)
```go
func (c *UserController) CreateUser(ctx context.Context, req *app.RequestContext) {
    // HTTP 请求处理
}
```

5. **注册路由** (`internal/interfaces/http/router.go`)
```go
v1.POST("/users", userController.CreateUser)
```

### API 端点

#### 健康检查
- `GET /health` - 服务健康状态
- `GET /ready` - 服务就绪状态

#### 认证相关
- `POST /api/v1/auth/login` - 用户登录
- `POST /api/v1/auth/refresh` - 刷新令牌
- `POST /api/v1/auth/logout` - 用户登出

#### 用户管理
- `GET /api/v1/users` - 获取用户列表
- `POST /api/v1/users` - 创建用户
- `GET /api/v1/users/:id` - 获取用户详情
- `PUT /api/v1/users/:id` - 更新用户
- `DELETE /api/v1/users/:id` - 删除用户

### 数据库迁移

运行数据库迁移：
```bash
go run cmd/migrate/main.go -e dev
```

### 运行测试

运行所有测试：
```bash
go test ./...
```

运行特定包的测试：
```bash
go test ./internal/domain/...
```

运行测试并生成覆盖率报告：
```bash
go test -cover ./...
```

### 代码生成

生成 Mock 文件：
```bash
go generate ./...
```

## Docker 支持

### 构建镜像

```bash
docker build -t go-template .
```

### 运行容器

```bash
docker run -p 8081:8081 -e APP_ENV=prod go-template
```

### Docker Compose

```yaml
version: '3.8'
services:
  app:
    build: .
    ports:
      - "8081:8081"
    environment:
      - APP_ENV=prod
    depends_on:
      - postgres
      - redis
  
  postgres:
    image: postgres:15
    environment:
      POSTGRES_DB: go_template
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
  
  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
```

## 部署

### 生产环境部署

1. **构建生产镜像**
```bash
docker build -t go-template:latest .
```

2. **运行生产容器**
```bash
docker run -d \
  --name go-template \
  -p 8081:8081 \
  -e APP_ENV=prod \
  go-template:latest
```

3. **使用 Docker Compose**
```bash
docker-compose up -d
```

### 环境变量

- `APP_ENV`: 运行环境 (dev/test/prod)
- `HTTP_PORT`: HTTP 服务端口
- `DB_HOST`: 数据库主机
- `DB_PORT`: 数据库端口
- `REDIS_HOST`: Redis 主机
- `REDIS_PORT`: Redis 端口

## 测试

### 单元测试

项目使用 Go 标准测试库和 Testify 进行单元测试：

```go
func TestUserService_CreateUser(t *testing.T) {
    // 测试实现
}
```

### 集成测试

集成测试使用真实的数据库连接：

```go
func TestUserRepository_Integration(t *testing.T) {
    // 集成测试实现
}
```

### Mock 测试

使用 GoMock 生成 Mock 对象：

```bash
mockgen -source=internal/domain/repository/user_repository.go -destination=test/mocks/user_repository_mock.go
```

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交改动 (`git commit -m 'Add amazing feature'`)
4. 推送分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

### 代码规范

- 遵循 Go 官方代码规范
- 使用 `gofmt` 格式化代码
- 使用 `golint` 检查代码质量
- 编写单元测试
- 添加必要的注释

## 许可证

本项目使用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 维护者

- [@lyonnee](https://github.com/lyonnee) - 项目创建者和主要维护者

## 致谢

- [CloudWeGo](https://github.com/cloudwego) - 提供高性能的 Hertz 框架
- [Uber](https://github.com/uber-go) - 提供优秀的 Zap 日志库
- [Spf13](https://github.com/spf13) - 提供强大的 Viper 配置库
- 所有为这个项目做出贡献的开发者