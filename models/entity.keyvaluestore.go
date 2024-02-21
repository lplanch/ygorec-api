package model

type KeyValueStore struct {
	Key   string `gorm:"type:VARCHAR(255) PRIMARY KEY"`
	Value string `gorm:"type:VARCHAR(255);NOT NULL"`
}
