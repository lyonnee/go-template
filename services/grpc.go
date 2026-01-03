package services

import (
	"context"
	"net"
	"time"

	"github.com/lyonnee/go-template/internal/infrastructure/config"
	grpcif "github.com/lyonnee/go-template/internal/interfaces/grpc"
	"github.com/lyonnee/go-template/pkg/di"
	"github.com/lyonnee/go-template/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func init() {
	s := NewGRPCService()
	RegisterService(s)
}

type GRPCService struct {
	srv  *grpc.Server
	lis  net.Listener
	addr string
}

func NewGRPCService() *GRPCService {
	conf := di.Get[config.Config]()
	return &GRPCService{
		addr: conf.Grpc.Port,
	}
}

func (s *GRPCService) Start() {
	// build gRPC server with options from config
	conf := di.Get[config.Config]()

	var opts []grpc.ServerOption
	if conf.Grpc.MaxConcurrentStreams > 0 {
		opts = append(opts, grpc.MaxConcurrentStreams(conf.Grpc.MaxConcurrentStreams))
	}

	// keepalive params
	ka := keepalive.ServerParameters{
		Time:                  conf.Grpc.Keepalive.Time,
		Timeout:               conf.Grpc.Keepalive.Timeout,
		MaxConnectionIdle:     conf.Grpc.Keepalive.MaxConnectionIdle,
		MaxConnectionAge:      conf.Grpc.Keepalive.MaxConnectionAge,
		MaxConnectionAgeGrace: conf.Grpc.Keepalive.MaxConnectionAgeGrace,
	}
	opts = append(opts, grpc.KeepaliveParams(ka))

	// keepalive enforcement policy
	ep := keepalive.EnforcementPolicy{
		MinTime:             conf.Grpc.Enforcement.MinTime,
		PermitWithoutStream: conf.Grpc.Enforcement.PermitWithoutStream,
	}
	opts = append(opts, grpc.KeepaliveEnforcementPolicy(ep))

	// interceptors (basic logging / recovery hooks)
	if conf.Grpc.Interceptors.EnableLogging {
		opts = append(opts, grpc.ChainUnaryInterceptor(loggingUnaryInterceptor()))
		opts = append(opts, grpc.ChainStreamInterceptor(loggingStreamInterceptor()))
	}
	if conf.Grpc.Interceptors.EnableRecovery {
		// 这里放一个简单的 panic recover 拦截器
		opts = append(opts, grpc.ChainUnaryInterceptor(recoveryUnaryInterceptor()))
		opts = append(opts, grpc.ChainStreamInterceptor(recoveryStreamInterceptor()))
	}

	s.srv = grpc.NewServer(opts...)

	// 注册各 gRPC 服务（默认 no-op，待代码生成后由带构建标签的文件实际注册）
	grpcif.RegisterGRPCServices(s.srv)

	// start listen
	var err error
	s.lis, err = net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal("failed to listen for gRPC")
		return
	}
	// serve in current goroutine; StartAll 已经为每个服务开 goroutine
	if err := s.srv.Serve(s.lis); err != nil {
		// Serve 只有在 Stop/GracefulStop 后或 fatal 错误才会返回
		log.Warn("gRPC server exited")
	}
}

func (s *GRPCService) Stop(ctx context.Context) {
	if s.srv == nil {
		return
	}
	// try graceful stop with context; fallback to Stop on timeout
	done := make(chan struct{})
	go func() {
		s.srv.GracefulStop()
		close(done)
	}()
	// compute remaining time from ctx
	timeout := 3 * time.Second
	if deadline, ok := ctx.Deadline(); ok {
		if left := time.Until(deadline); left < timeout {
			if left > 0 {
				timeout = left
			}
		}
	}
	select {
	case <-done:
		// graceful stop completed
	case <-time.After(timeout):
		log.Warn("gRPC graceful shutdown timed out, forcing stop")
		s.srv.Stop()
	}
	if s.lis != nil {
		_ = s.lis.Close()
	}
}

// --- Interceptors ---
func loggingUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		dur := time.Since(start)
		if err != nil {
			log.Warn("grpc unary call failed")
		} else {
			log.Info("grpc unary call completed")
		}
		_ = dur  // 可扩展为字段
		_ = info // 可扩展为方法名等
		return resp, err
	}
}

func loggingStreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()
		err := handler(srv, ss)
		dur := time.Since(start)
		if err != nil {
			log.Warn("grpc stream call failed")
		} else {
			log.Info("grpc stream call completed")
		}
		_ = dur
		_ = info
		return err
	}
}

func recoveryUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Error("grpc unary panic recovered")
				// 返回内部错误给客户端（可替换为自定义 status）
				err = grpc.Errorf(13, "internal") // 13 = Internal
			}
		}()
		return handler(ctx, req)
	}
}

func recoveryStreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Error("grpc stream panic recovered")
				err = grpc.Errorf(13, "internal")
			}
		}()
		return handler(srv, ss)
	}
}
