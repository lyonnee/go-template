package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/lyonnee/go-template/config"
	"github.com/lyonnee/go-template/pkg/log"
	"github.com/lyonnee/go-template/servers"
)

func main() {
	args := parseArgs()

	if err := config.Load(args.env); err != nil {
		// log.Fatalf("load config failed, err:%s", err)
	}

	if err := log.Initialize(config.Log()); err != nil {
		// log.Fatalf("init modules failed, err:%s", err)
	}

	go servers.StartHTTPServer(config.Http())
	go servers.StartRPCServer()

	log.Info("Server Running ...")

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info("Server Shutdown ...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Sync()
}

type args struct {
	env string
}

func parseArgs() *args {
	env := os.Getenv("APP_ENV")
	if env == "" {
		// Parse command line arguments
		args := os.Args[1:]
		for i := 0; i < len(args); i++ {
			switch args[i] {
			case "-e", "--env":
				if i+1 < len(args) {
					env = args[i+1]
					i++ // Skip the next argument since we consumed it
				}
				// Add more parameter cases here as needed
				// case "-p", "--password":
				//     if i+1 < len(args) {
				//         password = args[i+1]
				//         i++
				//     }
			}
		}
		// If still not set, use default
		if env == "" {
			env = "prod"
		}
	}

	return &args{
		env: env,
	}
}
