package model

import "gorm.io/gorm"

type MvTopCard struct {
	CardID     uint64  `gorm:"index:idx_card_banlist,unique,foreignKey:Card"`
	BanlistID  string  `gorm:"index:idx_card_banlist,unique,foreignKey:Banlist;DEFAULT NULL"`
	Percentage uint8   `gorm:"index:,INTEGER;NOT NULL;DEFAULT NULL"`
	Average    float32 `gorm:"DECIMAL(6,5);NOT NULL;DEFAULT NULL"`

	// Constraint
	Card    EntityCard    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Banlist EntityBanlist `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func AutoMigrateProcedureMvTopCards(db *gorm.DB) {
	db.Exec(`
		DROP PROCEDURE IF EXISTS refresh_mv_top_cards;
	`)
	db.Exec(`
		CREATE PROCEDURE refresh_mv_top_cards ()
		BEGIN
			TRUNCATE TABLE mv_top_cards;
			INSERT INTO mv_top_cards
			SELECT
				c.id AS card_id,
				NULL AS banlist_id,
				COUNT(DISTINCT g.deck_id) / (SELECT COUNT(*) FROM entity_decks) * 100 AS percentage,
				AVG(g.num_times_in_deck) AS average
			FROM entity_cards c
				JOIN (
					SELECT
						card_id,
						deck_id,
						COUNT(*) AS num_times_in_deck
					FROM
						graph_cards_belong_to_decks
					GROUP BY
						card_id,
						deck_id
					) g ON c.id = g.card_id
				JOIN entity_decks d ON g.deck_id = d.id
				GROUP BY c.id;
		END;
	`)
}
