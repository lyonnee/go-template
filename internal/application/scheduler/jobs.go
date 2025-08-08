package scheduler

import (
	"github.com/lyonnee/go-template/pkg/log"
	"github.com/robfig/cron/v3"
)

func RegisterJobs(s *cron.Cron) {
	s.AddFunc("0 * * * *", func() {
		// Example task: Log every hour
		log.Info("Hourly task executed")
	})
	s.AddFunc("0 0 * * *", func() {
		// Example task: Log every day at midnight
		log.Info("Daily task executed")
	})
}
