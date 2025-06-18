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
	"github.com/lyonnee/go-template/internal/infrastructure/di"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	"github.com/lyonnee/go-template/internal/infrastructure/persistence"
	"github.com/lyonnee/go-template/pkg/database"
	"github.com/lyonnee/go-template/pkg/id_generator"
	"github.com/lyonnee/go-template/pkg/logger"
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
	// Note: config must be loaded before any other service that depends on it
	conf, err := config.Load(env)
	if err != nil {
		stdLog.Printf("load config failed, err:%s", err)
		os.Exit(1)
	}

	di.AddSingleton[*config.Config](func() (*config.Config, error) {
		return conf, nil
	})

	// initialize ID generator
	// Note: ID generator must be initialized before any other service that depends on it
	if err := id_generator.Initialize(conf.App.NodeId); err != nil {
		stdLog.Printf("initialize ID generator failed, err:%s", err)
		os.Exit(1)
	}

	// initialize logger
	// Note: logger must be initialized before any other service that depends on it
	newLogger, err := logger.NewLogger(conf.Log)
	if err != nil {
		stdLog.Printf("init logger failed, err:%s", err)
		os.Exit(1)
	}
	di.AddSingleton[log.Logger](func() (log.Logger, error) {
		return newLogger, nil
	})

	// initialize database connection
	dbContext, err := database.NewDB(&conf.Persistence, newLogger)
	if err != nil {
		stdLog.Printf("init persistence failed, err:%s", err)
		os.Exit(1)
	}
	di.AddSingleton[persistence.DBContext](func() (persistence.DBContext, error) {
		return dbContext, nil
	})

	// bootstrap the application all services
	bootstrap.Initialize()
	bootstrap.Run()

	stdLog.Println("Server started successfully")
}

func shutdown() {
	stdLog.Println("Server is shutting down...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger := di.GetService[log.Logger]()
	logger.Sync()
}
