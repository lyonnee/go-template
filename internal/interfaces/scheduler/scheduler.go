package scheduler

import (
	"github.com/lyonnee/go-template/internal/interfaces/scheduler/jobs"
	"github.com/lyonnee/go-template/pkg/log"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	logger log.Logger
}

func (s *Scheduler) RegisterJobs(c *cron.Cron) {
	c.AddFunc("0 * * * *", func() {
		jobs.TestJob()
		// Example task: Log every hour
		s.logger.Info("Hourly task executed")
	})

	c.AddFunc("0 0 * * *", func() {
		// Example task: Log every day at midnight
		s.logger.Info("Daily task executed")
	})
}
