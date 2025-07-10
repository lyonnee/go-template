package main

import (
	"context"

	"os"
	"os/signal"
	"time"

	// for dependency injection
	_ "github.com/lyonnee/go-template/application/service"
	_ "github.com/lyonnee/go-template/domain/service"
	_ "github.com/lyonnee/go-template/infrastructure/repository_impl"
	_ "github.com/lyonnee/go-template/interfaces/http/controller"

	"github.com/lyonnee/go-template/infrastructure/database"
	"github.com/lyonnee/go-template/infrastructure/log"
	"github.com/lyonnee/go-template/interfaces/grpc"
	"github.com/lyonnee/go-template/interfaces/http"
)

func main() {
	go http.RunServer()
	go grpc.RunServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	http.CloseServer()
	database.Close()
	log.Sync()
}
