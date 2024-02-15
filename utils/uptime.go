package util

import (
	"time"
)

var startTime time.Time

func GetUptime() time.Duration {
	return time.Since(startTime)
}

func InitUptime() {
	startTime = time.Now()
}
