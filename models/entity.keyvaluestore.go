package model

type KeyValueStore struct {
	Key   string `gorm:"type:TEXT PRIMARY KEY"`
	Value string `gorm:"type:TEXT;NOT NULL"`
}
