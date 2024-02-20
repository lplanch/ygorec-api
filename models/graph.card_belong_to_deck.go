package model

type GraphCardsBelongToDecks struct {
	CardID   uint64 `gorm:"foreignKey:Card"`
	DeckID   string `gorm:"foreignKey:Deck"`
	Category uint8  `gorm:"type:INTEGER;NOT NULL;DEFAULT NULL"`

	// Constraint
	Card EntityCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Deck EntityDeck `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
