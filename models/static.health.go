package model

import "time"

type StaticHealth struct {
	Status string
	Uptime int64
	Date   time.Time
}
