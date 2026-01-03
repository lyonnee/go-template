package services

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lyonnee/go-template/internal/infrastructure/config"
	httpif "github.com/lyonnee/go-template/internal/interfaces/http"
	"github.com/lyonnee/go-template/pkg/di"
	"github.com/lyonnee/go-template/pkg/log"
)

func init() {
	s := NewHTTPService()
	RegisterService(s)
}

type HTTPService struct {
	engine *gin.Engine
	server *http.Server
}

func NewHTTPService() *HTTPService {
	conf := di.Get[config.Config]()

	// 设置gin模式
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	server := &http.Server{
		Addr:    conf.Http.Port,
		Handler: engine,
	}

	return &HTTPService{
		engine: engine,
		server: server,
	}
}

func (s *HTTPService) Start() {
	httpif.RegisterRoutes(s.engine)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("http server start failed: " + err.Error())
	}
}

func (s *HTTPService) Stop(ctx context.Context) {
	// Attempt graceful shutdown: stop accepting new requests and wait for in-flight requests
	cctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := s.server.Shutdown(cctx); err != nil {
		log.Warn("http graceful shutdown failed: " + err.Error())
	}
}
