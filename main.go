package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lyonnee/go-template/internal/infrastructure/database"
	"github.com/lyonnee/go-template/pkg/log"
	"github.com/lyonnee/go-template/services"

	_ "github.com/lyonnee/go-template/internal/infrastructure/repository_impl"
)

func main() {
	services.StartAll()

	// Wait for termination signal (SIGINT/SIGTERM)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Info("received shutdown signal, start graceful shutdown")

	// Begin graceful shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		// Stop services and close resources
		services.StopAll(shutdownCtx)
		_ = database.CloseWithContext(shutdownCtx)
		close(done)
	}()

	select {
	case <-done:
		// graceful shutdown completed
		log.Info("graceful shutdown completed")
	case <-shutdownCtx.Done():
		// timeout reached; proceed with forced exit
		log.Warn("graceful shutdown timed out")
	}

	log.Sync()
}
