package model

type GraphCardsBelongToDecks struct {
	CardID   uint64 `gorm:"index:,foreignKey:Card"`
	DeckID   string `gorm:"index:,foreignKey:Deck"`
	Category uint8  `gorm:"type:INTEGER;NOT NULL;DEFAULT NULL"`

	// Constraint
	Card EntityCard `gorm:"foreignKey:CardID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Deck EntityDeck `gorm:"foreignKey:DeckID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
