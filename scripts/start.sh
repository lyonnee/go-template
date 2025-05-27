#!/bin/bash

# 启动 Go 模板应用程序的脚本

set -e

echo "🚀 启动 Go Template 应用程序..."

# 检查环境变量
APP_ENV=${APP_ENV:-dev}
echo "📝 使用环境: $APP_ENV"

# 构建应用程序
echo "🔨 构建应用程序..."
go build -o bin/server ./cmd/server

# 启动应用程序
echo "▶️  启动服务器..."
./bin/server -e $APP_ENV
