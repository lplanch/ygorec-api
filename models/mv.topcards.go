package model

import "gorm.io/gorm"

type MvTopCard struct {
	CardID    uint64  `gorm:"index:idx_card_banlist,unique"`
	BanlistID string  `gorm:"index:idx_card_banlist,unique;DEFAULT NULL"`
	Amount    uint32  `gorm:"index:INT UNSIGNED;NOT NULL;DEFAULT NULL"`
	Average   float32 `gorm:"DECIMAL(6,5);NOT NULL;DEFAULT NULL"`

	// Constraint
	Card    EntityCard    `gorm:"foreignKey:CardID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Banlist EntityBanlist `gorm:"foreignKey:BanlistID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func AutoMigrateProcedureMvTopCards(db *gorm.DB) {
	db.Exec(`
		DROP PROCEDURE IF EXISTS refresh_mv_top_cards;
	`)
	db.Exec(`
		CREATE PROCEDURE refresh_mv_top_cards (IN banlist_id VARCHAR(255))
		BEGIN
			IF banlist_id IS NULL THEN
				DELETE FROM mv_top_cards mv WHERE mv.banlist_id IS NULL;
				INSERT INTO mv_top_cards
				SELECT
					s.card_id AS card_id,
					NULL AS banlist_id,
					COUNT(*) AS amount,
					AVG(s.amount) AS average
				FROM (
					SELECT
						g.card_id,
						g.deck_id,
						COUNT(*) AS amount
					FROM graph_cards_belong_to_decks g
					JOIN entity_decks d ON g.deck_id = d.id
					GROUP BY g.card_id, g.deck_id
				) s
				GROUP BY s.card_id;
			ELSE
				DELETE FROM mv_top_cards mv WHERE mv.banlist_id = banlist_id;
				INSERT INTO mv_top_cards
				SELECT
					s.card_id AS card_id,
					banlist_id,
					COUNT(*) AS amount,
					AVG(s.amount) AS average
				FROM (
					SELECT
						g.card_id,
						g.deck_id,
						COUNT(*) AS amount
					FROM graph_cards_belong_to_decks g
					JOIN entity_decks d ON g.deck_id = d.id AND d.updated_at >= banlist_id
					GROUP BY g.card_id, g.deck_id
				) s
				GROUP BY s.card_id;
			END IF;
		END;
	`)
}
