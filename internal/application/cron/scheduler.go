package cron

import (
	"context"

	"github.com/lyonnee/go-template/pkg/di"
	"github.com/robfig/cron/v3"
)

func init() {
	di.AddSingleton[*Scheduler](NewScheduler)
}

func NewScheduler() (*Scheduler, error) {
	return &Scheduler{
		cron: cron.New(cron.WithSeconds()),
	}, nil
}

// 用于启动调度器
func RunScheduler() {
	s := di.Get[*Scheduler]()

	registerTasks(s)

	s.Run()
}

type Scheduler struct {
	cron *cron.Cron
}

func (s *Scheduler) AddJob(spec string, job func(context.Context)) (cron.EntryID, error) {
	return s.cron.AddFunc(spec, func() { job(context.Background()) })
}

func (s *Scheduler) Run() {
	s.cron.Start()
}

func (s *Scheduler) Shutdown(ctx context.Context) {
	s.cron.Stop()
}
