package model

import "time"

type StaticVersion struct {
	CardsLastCommit string
	EnumLastCommit  string
	LastUpdate      time.Time
	CardsAmount     uint32
}
