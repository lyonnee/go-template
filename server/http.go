package server

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/internal/interfaces/http"
)

func StartHTTPServer(conf config.HttpConfig) {
	s := server.New(
		server.WithHostPorts(conf.Port),
	)

	http.Register(s)

	s.Spin()
}
