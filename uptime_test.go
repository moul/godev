package godev

import (
	"fmt"
	"time"
)

func ExampleUptime() {
	uptime := Uptime()
	if uptime > 1*time.Nanosecond && uptime < 1*time.Minute {
		fmt.Println("uptime is less than a minute")
	}
	// Output: uptime is less than a minute
}
