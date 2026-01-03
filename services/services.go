package services

import (
	"context"

	"github.com/lyonnee/go-template/pkg/log"
)

type Service interface {
	Start()
	Stop(ctx context.Context)
}

var (
	services []Service
)

func RegisterService(s Service) {
	services = append(services, s)
}

// StartAll 启动所有服务
func StartAll() {
	log.Info("starting all services")
	for _, s := range services {
		go s.Start()
	}
}

// StopAll 停止所有服务（需要 context）
func StopAll(ctx context.Context) {
	log.Info("stopping all services")
	// stop in reverse order to respect dependencies
	for i := len(services) - 1; i >= 0; i-- {
		s := services[i]
		s.Stop(ctx)
	}
	log.Info("all services stopped")
}
