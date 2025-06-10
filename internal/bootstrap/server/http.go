package server

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/internal/infrastructure/di"
	"github.com/lyonnee/go-template/internal/interfaces/http"
)

func StartHTTPServer() {
	config := di.GetService[*config.Config]().Http

	s := server.New(
		server.WithHostPorts(config.Port),
	)

	http.Register(s)

	s.Spin()
}
