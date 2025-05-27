package main

import (
	"context"
	"flag"
	stdLog "log"
	"os"
	"os/signal"
	"time"

	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/internal/infrastructure/log"
	pkgLog "github.com/lyonnee/go-template/pkg/log"
	"github.com/lyonnee/go-template/pkg/persistence"
	"github.com/lyonnee/go-template/server"
)

func main() {
	var (
		env = flag.String("env", "dev", "Environment (dev, test, prod)")
	)
	flag.Parse()

	if err := config.Load(*env); err != nil {
		stdLog.Printf("load config failed, err:%s", err)
		os.Exit(1)
	}

	zapLogger, err := pkgLog.NewZapLogger(config.Log())
	if err != nil {
		stdLog.Printf("init zap logger failed, err:%s", err)
		os.Exit(1)
	}
	logger := pkgLog.NewZapSugarLogger(zapLogger)
	log.SetGLogger(logger)

	if err := persistence.Initialize(config.Persistence(), logger); err != nil {
		stdLog.Printf("init persistence failed, err:%s", err)
		os.Exit(1)
	}

	go server.StartHTTPServer(config.Http())
	go server.StartRPCServer()

	logger.Info("Server Running ...")

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logger.Info("Server Shutdown ...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Sync()
}

type args struct {
	env string
}
