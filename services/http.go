package services

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/lyonnee/go-template/internal/infrastructure/config"
	"github.com/lyonnee/go-template/internal/interfaces/http"
	"github.com/lyonnee/go-template/pkg/di"
)

var s *server.Hertz

func StartHTTPServer() {
	conf := di.Get[config.Config]()

	s = server.New(
		server.WithHostPorts(conf.Http.Port),
	)

	http.RegisterRoutes(s)

	s.Spin()
}

func StopHTTPServer() {
	if s != nil {
		s.Close()
	}
}
