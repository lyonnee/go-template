package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/lyonnee/go-template/internal/infrastructure/database"
	"github.com/lyonnee/go-template/pkg/log"
	"github.com/lyonnee/go-template/services"

	_ "github.com/lyonnee/go-template/internal/infrastructure/repository_impl"
)

func main() {
	services.StartAllServices()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	services.StopAllServices()
	database.Close()
	log.Sync()
}
