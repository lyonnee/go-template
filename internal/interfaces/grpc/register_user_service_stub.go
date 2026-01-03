//go:build !userpb

package grpcif

import grpc "google.golang.org/grpc"

// 默认情况下（未生成 proto 代码）不注册任何服务，避免编译失败。
// 生成代码后会用另一个文件（带构建标签）覆盖该实现进行实际注册。
func RegisterGRPCServices(s *grpc.Server) {}
