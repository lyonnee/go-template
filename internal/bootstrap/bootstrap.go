package bootstrap

import "github.com/lyonnee/go-template/internal/bootstrap/server"

// Run initializes the bootstrap process for the application.
func Run() {
	registerServices()

	go server.StartHTTPServer()
	go server.StartRPCServer()
}
