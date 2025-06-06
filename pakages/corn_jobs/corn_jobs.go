package corn_jobs

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

var CronInstance *cron.Cron

func Config() {
	// Create a new cron scheduler
	c := cron.New(cron.WithSeconds())

	// Schedule a task to run every 10 seconds
	_, err := c.AddFunc("*/10 * * * * *", func() {
		fmt.Println("Scheduled task executed at", time.Now())
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	c.Start()        // start the scheduler
	CronInstance = c // Store for later if needed
}
