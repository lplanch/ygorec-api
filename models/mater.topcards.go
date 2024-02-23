package model

type MvTopCard struct {
	CardID     uint64  `gorm:"index:idx_card_banlist,unique,foreignKey:Card"`
	BanlistID  string  `gorm:"index:idx_card_banlist,unique,foreignKey:Banlist;DEFAULT NULL"`
	Percentage uint8   `gorm:"index:,INTEGER;NOT NULL;DEFAULT NULL"`
	Average    float32 `gorm:"DECIMAL(6,5);NOT NULL;DEFAULT NULL"`

	// Constraint
	Card    EntityCard    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Banlist EntityBanlist `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// func AutoMigrateRefreshMvTopCards(db *gorm.DB) {
// 	db.Exec(`
// 	DROP PROCEDURE IF EXISTS refresh_mv_top_cards;

// 	DELIMITER $$
// 	CREATE PROCEDURE refresh_mv_top_cards (
// 	    OUT rc INT
// 	)
// 	BEGIN
// 		TRUNCATE TABLE mv_top_cards;
// 		INSERT INTO mv_top_cards
// 		SELECT
// 			id as card_id,

// 			COUNT(*)
// 		FROM entity_cards
// 			GROUP BY id;
// 		SET rc = 0;
// 	END;
// 	$$

// 	DELIMITER ;
// 	`)
// }
