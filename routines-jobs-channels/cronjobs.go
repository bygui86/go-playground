package main

import (
	"fmt"
	"time"
	
	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Println("*** CronJob tests started")
	
	fmt.Println("Build Cron")
	cronJob := cron.New(
		// cron.WithLocation(time.UTC),
		cron.WithSeconds(),
	)
	
	fmt.Println("Add functions to Cron")
	// second | minute | hour | day-of-month | month | day-of-week
	cronJob.AddFunc("*/5 * * * * *", execution)
	// cronJob.AddFunc("1 * * * * *", execution)
	// cronJob.AddFunc("1 * * * * *", execution)
	
	fmt.Println("First Cron start")
	cronJob.Start()
	// for _, job := range cronJob.Entries() {
	// 	fmt.Printf("cron entryId %d\n", job.ID)
	// 	fmt.Printf("cron next %v\n", job.Next)
	// 	fmt.Printf("cron valid %t\n", job.Valid())
	// }
	time.Sleep(20 *time.Second)
	fmt.Println("First Cron stop")
	cronJob.Stop()
	
	time.Sleep(20 *time.Second)
	
	fmt.Println("Second Cron start")
	cronJob.Start()
	time.Sleep(20 *time.Second)
	fmt.Println("Second Cron stop")
	cronJob.Stop()
	
	fmt.Println("*** CronJob tests completed")
}

func execution() {
	now := time.Now()
	fmt.Printf("this is a job execution at %v\n", now)
}
