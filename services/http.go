package services

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/lyonnee/go-template/internal/infrastructure/config"
	"github.com/lyonnee/go-template/internal/interfaces/http"
	"github.com/lyonnee/go-template/pkg/di"
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

func (s *HTTPService) Stop() {
	s.h.Close()
}
