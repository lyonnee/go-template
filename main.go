package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/lyonnee/go-template/internal/application/cron"
	"github.com/lyonnee/go-template/internal/infrastructure/database"
	"github.com/lyonnee/go-template/internal/interfaces/grpc"
	"github.com/lyonnee/go-template/internal/interfaces/http"
	"github.com/lyonnee/go-template/pkg/log"

	_ "github.com/lyonnee/go-template/internal/infrastructure/repository_impl"
)

func main() {
	go cron.RunScheduler()
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
