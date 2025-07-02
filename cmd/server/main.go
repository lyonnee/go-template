package main

import (
	"context"
	"flag"

	stdLog "log"
	"os"
	"os/signal"
	"time"

	"github.com/lyonnee/go-template/bootstrap"
	"github.com/lyonnee/go-template/bootstrap/di"
	"go.uber.org/zap"
)

func main() {
	var (
		env = flag.String("env", "dev", "Environment (dev, test, prod)")
	)
	flag.Parse()

	if err := bootstrap.Initialize(*env); err != nil {
		stdLog.Printf("bootstrap failed, err:%s", err)
		os.Exit(1)
	}

	// wait for interrupt signal to gracefully shutdown the server (with 5 seconds timeout)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	shutdown()
}

func shutdown() {
	stdLog.Println("Server is shutting down...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger := di.Get[*zap.Logger]()
	logger.Sync()
}
