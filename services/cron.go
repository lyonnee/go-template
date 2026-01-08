package services

import (
	"context"
	"time"

	"github.com/lyonnee/go-template/internal/interfaces/scheduler"
	"github.com/lyonnee/go-template/pkg/log"
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

func (s *CronService) Stop(ctx context.Context) {
	// Stop() returns a context that is closed when jobs complete
	doneCtx := s.c.Stop()
	timeout := 3 * time.Second
	if deadline, ok := ctx.Deadline(); ok {
		if left := time.Until(deadline); left < timeout {
			timeout = left
		}
	}
	select {
	case <-doneCtx.Done():
		// graceful stop completed
	case <-time.After(timeout):
		// timeout waiting for jobs; proceed
		log.Warn("cron shutdown timed out")
	}
}
