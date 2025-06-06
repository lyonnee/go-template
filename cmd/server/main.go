package main

import (
	"context"
	"flag"

	stdLog "log"
	"os"
	"os/signal"
	"time"

	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/internal/bootstrap"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/lyonnee/go-template/pkg"
	"github.com/lyonnee/go-template/pkg/container"
)

func main() {
	var (
		env = flag.String("env", "dev", "Environment (dev, test, prod)")
	)
	flag.Parse()

	start(*env)

	// wait for interrupt signal to gracefully shutdown the server (with 5 seconds timeout)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	shutdown()
}

func start(env string) {
	// initialize config
	if err := config.Load(env); err != nil {
		stdLog.Printf("load config failed, err:%s", err)
		os.Exit(1)
	}

	// initialize pkg modules
	if err := pkg.Initialize(); err != nil {
		stdLog.Printf("init pkg modules failed, err:%s", err)
		os.Exit(1)
	}

	// bootstrap the application all services
	bootstrap.Run()

	stdLog.Println("Server started successfully")
}

func shutdown() {
	stdLog.Println("Server is shutting down...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger := container.GetService[log.Logger]()
	logger.Sync()
}
