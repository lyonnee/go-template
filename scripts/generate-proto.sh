#!/bin/bash

# Proto文件生成脚本
set -e

echo "🔨 Generating protobuf code..."

# 检查protoc是否安装
if ! command -v protoc &> /dev/null; then
    echo "❌ protoc is not installed. Please install it first."
    exit 1
fi

# 获取GOPATH
GOPATH=$(go env GOPATH)

# 检查protoc插件是否安装
if [ ! -f "$GOPATH/bin/protoc-gen-go" ]; then
    echo "📦 Installing protoc-gen-go..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
fi

if [ ! -f "$GOPATH/bin/protoc-gen-go-grpc" ]; then
    echo "📦 Installing protoc-gen-go-grpc..."
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
fi

# 生成proto代码
echo "🚀 Generating Go code from proto files..."

# 确保userpb目录存在
mkdir -p internal/interfaces/grpc/userpb

# 生成user.proto
protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    --plugin=protoc-gen-go="$GOPATH/bin/protoc-gen-go" \
    --plugin=protoc-gen-go-grpc="$GOPATH/bin/protoc-gen-go-grpc" \
    internal/interfaces/grpc/user.proto

# 移动生成的文件到userpb目录
mv internal/interfaces/grpc/user.pb.go internal/interfaces/grpc/userpb/
mv internal/interfaces/grpc/user_grpc.pb.go internal/interfaces/grpc/userpb/

echo "✅ Proto code generation completed successfully!"
echo ""
echo "Generated files:"
ls -lh internal/interfaces/grpc/userpb/*.pb.go
