package model

import "time"

type StaticVersion struct {
	LastCommit string
	LastUpdate time.Time
}
