package server

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/internal/interfaces/http"
	"github.com/lyonnee/go-template/pkg/container"
)

func StartHTTPServer() {
	config := container.GetService[*config.Config]().Http

	s := server.New(
		server.WithHostPorts(config.Port),
	)

	http.Register(s)

	s.Spin()
}
