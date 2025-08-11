package services

import (
	"time"

	"github.com/lyonnee/go-template/internal/application/scheduler"
	"github.com/robfig/cron/v3"
)

var c *cron.Cron

func StartCronScheduler() {
	c = cron.New(cron.WithSeconds(), cron.WithLocation(time.UTC))

	scheduler.RegisterScheduledJobs(c)

	c.Start()
}

func StopCronScheduler() {
	if c != nil {
		c.Stop()
	}
}
