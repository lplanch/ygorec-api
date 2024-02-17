package model

type KeyValueStore struct {
	Key   string `gorm:"type:TEXT PRIMARYKEY"`
	Value string `gorm:"type:TEXT"`
}
