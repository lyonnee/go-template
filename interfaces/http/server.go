package http

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/lyonnee/go-template/infrastructure/config"
	"github.com/lyonnee/go-template/infrastructure/di"
)

var s *server.Hertz

func RunServer() {
	conf := di.Get[config.Config]()

	s = server.New(
		server.WithHostPorts(conf.Http.Port),
	)

	registerRoutes(s)

	s.Spin()
}

func CloseServer() {
	if s != nil {
		s.Close()
	}
}
