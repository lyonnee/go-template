package http

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/lyonnee/go-template/config"
)

func RunServer(conf config.HttpConfig) {
	s := server.New(
		server.WithHostPorts(conf.Port),
	)

	registerRoutes(s)

	s.Spin()
}
