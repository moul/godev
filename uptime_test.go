package godev_test

import (
	"fmt"
	"time"

	"moul.io/godev"
)

func ExampleUptime() {
	uptime := godev.Uptime()
	if uptime > 1*time.Nanosecond && uptime < 1*time.Minute {
		fmt.Println("uptime is less than a minute")
	}
	// Output: uptime is less than a minute
}
