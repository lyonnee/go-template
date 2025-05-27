# 多阶段构建
FROM golang:1.23.7-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装必要的包
RUN apk add --no-cache git ca-certificates tzdata

# 设置环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 复制 go mod 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN go build -ldflags="-w -s" -o server ./cmd/server/main.go
RUN go build -ldflags="-w -s" -o migrate ./cmd/migrate/main.go

# 生产阶段
FROM alpine:latest

# 设置时区
RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Shanghai

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/server .
COPY --from=builder /app/migrate .

# 复制配置文件
COPY config*.yaml ./
COPY sql/ ./sql/

# 创建日志目录
RUN mkdir -p _logs

# 暴露端口
EXPOSE 8080

# 设置用户
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup
USER appuser

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/v1/health || exit 1

# 启动应用
CMD ["./server", "--env=prod"]