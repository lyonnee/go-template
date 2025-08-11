package scheduler

import (
	"github.com/lyonnee/go-template/internal/application/scheduler/jobs"
	"github.com/lyonnee/go-template/pkg/log"
	"github.com/robfig/cron/v3"
)

func RegisterScheduledJobs(s *cron.Cron) {
	s.AddFunc("0 * * * *", func() {
		jobs.TestJob()
		// Example task: Log every hour
		log.Info("Hourly task executed")
	})
	s.AddFunc("0 0 * * *", func() {
		// Example task: Log every day at midnight
		log.Info("Daily task executed")
	})
}
