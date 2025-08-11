#!/bin/bash

# 构建脚本
set -e

echo "🚀 Building Go Template Project..."

# 检查Go版本
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo "Go version: $GO_VERSION"

# 设置环境变量
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

# 清理之前的构建
echo "🧹 Cleaning previous builds..."
rm -rf ./bin
mkdir -p ./bin

# 下载依赖
echo "📦 Downloading dependencies..."
go mod download
go mod tidy

# 运行测试
echo "🧪 Running tests..."
go test ./... -v

# 构建应用
echo "🔨 Building application..."
go build -ldflags="-w -s" -o ./bin/server ./cmd/server/main.go

# 检查构建结果
if [ -f "./bin/server" ]; then
    echo "✅ Server build successful"
    ls -lh ./bin/server
else
    echo "❌ Server build failed"
    exit 1
fi

if [ -f "./bin/migrate" ]; then
    echo "✅ Migrate build successful"
    ls -lh ./bin/migrate
else
    echo "❌ Migrate build failed"
    exit 1
fi

echo "🎉 Build completed successfully!"
