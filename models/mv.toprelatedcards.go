package model

import "gorm.io/gorm"

type MvTopRelatedCards struct {
	FromCardID uint64 `gorm:"index:idx_top_related_card_card,unique"`
	ToCardID   uint64 `gorm:"index:idx_top_related_card_card,unique"`
	DeckAmount uint32 `gorm:"index:INT UNSIGNED;NOT NULL;DEFAULT NULL"`
	CardAmount uint32 `gorm:"index:INT UNSIGNED;NOT NULL;DEFAULT NULL"`

	// Constraint
	FromCard EntityCard `gorm:"foreignKey:FromCardID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ToCard   EntityCard `gorm:"foreignKey:ToCardID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func AutoMigrateProcedureMvTopRelatedCards(db *gorm.DB) {
	db.Exec(`
		DROP PROCEDURE IF EXISTS refresh_mv_top_related_cards;
	`)
	db.Exec(`
		CREATE PROCEDURE refresh_mv_top_related_cards (IN p_card_id INT UNSIGNED)
		BEGIN
			DELETE FROM mv_top_related_cards mv WHERE mv.from_card_id = card_id;
			INSERT INTO mv_top_related_cards
			SELECT
				p_card_id AS from_card_id,
				a.card_id AS to_card_id,
				COUNT(a.amount) AS deck_amount,
				SUM(a.amount) AS card_amount
			FROM (
				SELECT
					(CASE WHEN ISNULL(c.id) THEN g.card_id ELSE c.alias END) AS card_id,
					COUNT(*) AS amount
				FROM graph_cards_belong_to_decks g
				    LEFT OUTER JOIN entity_cards c ON c.id = g.card_id AND c.alias != 0
					WHERE card_id != p_card_id AND EXISTS(
						SELECT c_d.deck_id FROM graph_cards_belong_to_decks c_d
						LEFT OUTER JOIN entity_cards ec ON ec.alias = c_d.card_id
						WHERE (c_d.card_id = p_card_id OR ec.id = p_card_id) AND c_d.deck_id = g.deck_id
						GROUP BY c_d.deck_id, c_d.card_id
					)
					GROUP BY card_id, g.deck_id
			) a
				GROUP BY a.card_id
				ORDER BY deck_amount DESC, card_amount DESC, a.card_id ASC
				LIMIT 200;
		END;
	`)
}
