package model

type GraphCardsBelongToBanlists struct {
	CardID    uint64 `gorm:"index:idx_graph_card_banlist,unique,foreignKey:Card"`
	BanlistID string `gorm:"index:idx_graph_card_banlist,unique,foreignKey:Banlist"`
	Status    uint8  `gorm:"type:TINYINT UNSIGNED;NOT NULL"`

	// Constraint
	Card    EntityCard    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Banlist EntityBanlist `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
