package model

type MvDeckArchetypes struct {
	DeckID       string `gorm:"index:idx_graph_deck_archetype,unique,foreignKey:Deck"`
	ArechetypeID uint64 `gorm:"index:idx_graph_deck_archetype,unique,foreignKey:Archetype"`
	Weight       uint8  `gorm:"type:TINYINT UNSIGNED;NOT NULL"`

	// Constraint
	Deck      EntityDeck    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Archetype EnumArchetype `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
