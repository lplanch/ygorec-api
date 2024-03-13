package model

type GraphCardsBelongToBanlists struct {
	CardID    uint64 `gorm:"index:idx_graph_card_banlist,unique"`
	BanlistID string `gorm:"index:idx_graph_card_banlist,unique"`
	Status    uint8  `gorm:"type:TINYINT UNSIGNED;NOT NULL"`

	// Constraint
	Card    EntityCard    `gorm:"foreignKey:CardID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Banlist EntityBanlist `gorm:"foreignKey:BanlistID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
