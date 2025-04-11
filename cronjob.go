package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func Cronjob() {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("*/3 * * * * *", func() {
		fmt.Println("Hello World! every 3 second")
	})

	c.Start()

	select {}
}
