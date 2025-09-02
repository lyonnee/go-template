package services

import (
	"time"

	"github.com/lyonnee/go-template/internal/application/scheduler"
	"github.com/robfig/cron/v3"
)

func init() {
	s := NewCronService()
	RegisterService(s)
}

type CronService struct {
	c *cron.Cron
}

func NewCronService() *CronService {
	return &CronService{
		c: cron.New(cron.WithSeconds(), cron.WithLocation(time.UTC)),
	}
}

func (s *CronService) Start() {
	scheduler.RegisterScheduledJobs(s.c)
	s.c.Start()
}

func (s *CronService) Stop() {
	s.c.Stop()
}
