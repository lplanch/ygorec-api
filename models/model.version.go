package model

import "time"

type StaticVersion struct {
	CardsLastCommit string
	EnumsLastCommit string
	CardsLastUpdate time.Time
	EnumsLastUpdate time.Time
}
