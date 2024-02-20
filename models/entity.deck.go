package model

import (
	"time"

	"gorm.io/gorm"
)

type EntityDeck struct {
	ID        string `gorm:"type:TEXT PRIMARY KEY"`
	Source    string `gorm:"type:TEXT"`
	DeckType  string `gorm:"type:TEXT"`
	UpdatedAt time.Time
}

func (entity *EntityDeck) BeforeCreate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now()
	return nil
}

func (entity *EntityDeck) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now()
	return nil
}
