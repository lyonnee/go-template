package services

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/lyonnee/go-template/internal/infrastructure/config"
	"github.com/lyonnee/go-template/internal/interfaces/http"
	"github.com/lyonnee/go-template/pkg/di"
	"github.com/lyonnee/go-template/pkg/log"
)

func init() {
	s := NewHTTPService()
	RegisterService(s)
}

type HTTPService struct {
	h *server.Hertz
}

func NewHTTPService() *HTTPService {
	conf := di.Get[config.Config]()

	s := server.New(
		server.WithHostPorts(conf.Http.Port),
	)
	return &HTTPService{
		h: s,
	}
}

func (s *HTTPService) Start() {
	http.RegisterRoutes(s.h)
	s.h.Spin()
}

func (s *HTTPService) Stop(ctx context.Context) {
	// Attempt graceful shutdown: stop accepting new requests and wait for in-flight requests
	cctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.h.Shutdown(cctx); err != nil {
		// Fallback to force close on timeout or error
		log.Warn("http graceful shutdown failed, forcing close")
		s.h.Close()
	}
}
