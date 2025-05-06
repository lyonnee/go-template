# Go Template

一个基于 Go 语言的现代化 Web 服务项目模板，采用领域驱动设计(DDD)和整洁架构(Clean Architecture)思想构建。

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

- 📦 整洁架构 (Clean Architecture)
- 🎯 领域驱动设计 (DDD)
- 🔐 JWT 认证
- 📝 结构化日志 (Zap)
- 🗄️ PostgreSQL 数据库
- 🔄 优雅关机
- 🐳 Docker 支持
- 📊 Prometheus 指标
- ⚡ 高性能 HTTP 路由 (Fiber)
- 🔌 gRPC 支持
- 📨 消息队列集成
- 💾 缓存支持
- ⏰ 定时任务
- 🔧 多环境配置

## 项目结构

```
go-template/                          # 项目根目录
├── cmd/                             # 命令行入口目录
│   ├── migrate/                     # 数据库迁移工具
│   ├── scheduler/                   # 定时任务调度器
│   └── server/                      # 主服务器
│
├── sql/                             # sql描述文件
│
├── config/                          # 配置管理模块
│   ├── config.go                    # 配置加载逻辑
│   └── model.go                     # 配置结构定义
│
├── internal/                        # 内部应用代码
│   ├── application/                 # 应用层：处理业务流程
│   │   ├── auth_service.go         # 认证服务接口
│   │   └── impl/                   # 服务实现
│   ├── domain/                     # 领域层：核心业务逻辑
│   ├── infrastructure/            # 基础设施适配层
│   │   └── repositories/         # 仓储实现
│   └── interfaces/               # 接口层
│       ├── grpc/                # gRPC 接口
│       └── http/                # HTTP 接口
│           ├── controller/      # 控制器
│           ├── dto/            # 数据传输对象
│           └── middleware/     # 中间件
│
├── pkg/                        # 公共包
│   ├── modules/               # 基础模块
│   │   ├── auth/            # 认证模块
│   │   ├── blockchain/     # 区块链集成
│   │   ├── cache/         # 缓存模块
│   │   ├── logger/        # 日志模块
│   │   ├── mq/           # 消息队列
│   │   └── persistence/  # 数据持久化
│   └── scheduler/         # 调度器模块
│
└── servers/               # 服务器实现
    ├── http.go           # HTTP 服务器
    └── rpc.go            # RPC 服务器

```

## 架构说明

项目采用整洁架构（Clean Architecture）和领域驱动设计（DDD）思想构建，分为以下几层：

1. **接口层** (Interface Layer)
   - 处理 HTTP 和 gRPC 请求
   - 请求参数验证
   - 响应封装

2. **应用层** (Application Layer)
   - 处理业务流程
   - 编排领域对象
   - 事务管理

3. **领域层** (Domain Layer)
   - 核心业务逻辑
   - 领域模型
   - 领域服务

4. **基础设施层** (Infrastructure Layer)
   - 数据持久化
   - 消息队列
   - 缓存
   - 日志
   - 认证

## 开发指南

### 添加新功能

1. 在 `internal/domain` 中定义领域模型
2. 在 `internal/application` 中实现业务逻辑
3. 在 `internal/interfaces` 中添加 API 接口
4. 在 `internal/infrastructure` 中实现基础设施

### 配置管理

配置文件位于项目根目录：
- `config.dev.yaml`: 开发环境
- `config.test.yaml`: 测试环境
- `config.prod.yaml`: 生产环境

### 数据库迁移

```bash
go run cmd/migrate/main.go
```

### 运行测试

```bash
go test ./...
```

## Docker 支持

构建镜像：
```bash
docker build -t your-app-name .
```

运行容器：
```bash
docker run -p 8080:8080 your-app-name
```

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交改动 (`git commit -m 'Add amazing feature'`)
4. 推送分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 许可证

本项目使用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 维护者

- [@lyonnee](https://github.com/lyonnee)

## 致谢

感谢所有为这个项目做出贡献的开发者。