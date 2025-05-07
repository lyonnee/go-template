package server

import (
	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/internal/interfaces/http"
)

func StartHTTPServer(conf config.HttpConfig) {
	app := http.New()

	app.Listen(conf.Port)
}
