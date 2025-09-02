package services

func init() {
	s := NewGRPCService()
	RegisterService(s)
}

type GRPCService struct {
}

func NewGRPCService() *GRPCService {
	return &GRPCService{}
}

func (s *GRPCService) Start() {
	// 启动 gRPC 服务
}

func (s *GRPCService) Stop() {
	// 停止 gRPC 服务
}
