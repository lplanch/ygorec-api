package model

import "gorm.io/gorm"

type MvTopArchetypeCard struct {
	ArchetypeID uint64 `gorm:"index:idx_top_archetype_archetype_card,unique"`
	CardID      uint64 `gorm:"index:idx_top_archetype_archetype_card,unique"`
	DeckAmount  uint32 `gorm:"index:INT UNSIGNED;NOT NULL;DEFAULT NULL"`
	CardAmount  uint32 `gorm:"index:INT UNSIGNED;NOT NULL;DEFAULT NULL"`

	// Constraint
	Archetype EnumArchetype `gorm:"foreignKey:ArchetypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Card      EntityCard    `gorm:"foreignKey:CardID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func AutoMigrateProcedureMvTopArchetypeCards(db *gorm.DB) {
	db.Exec(`
		DROP PROCEDURE IF EXISTS refresh_mv_top_archetype_cards;
	`)
	db.Exec(`
		CREATE PROCEDURE refresh_mv_top_archetype_cards (IN archetype_id INT UNSIGNED)
		BEGIN
			DELETE FROM mv_top_archetype_cards mv WHERE mv.archetype_id = archetype_id;
			INSERT INTO mv_top_archetype_cards
			SELECT
				b.archetype_id,
				b.card_id,
				COUNT(b.CardAmount) AS deck_amount,
				SUM(b.CardAmount) AS card_amount
			FROM
				(
					SELECT
						da.archetype_id,
						da.deck_id,
						c_d.card_id,
						COUNT(c_d.card_id) AS CardAmount
					FROM
						mv_deck_archetypes da
						LEFT JOIN graph_cards_belong_to_decks c_d ON c_d.deck_id = da.deck_id
					WHERE
						da.archetype_id = archetype_id
						AND da.weight > 5
					GROUP BY
						da.deck_id,
						c_d.card_id,
						da.archetype_id
				) b
			JOIN entity_cards ec ON ec.id = b.card_id
			WHERE match_archetype(ec.set_code, archetype_id)
			GROUP BY b.card_id, b.archetype_id;
			INSERT INTO mv_top_archetype_cards
			SELECT
				b.archetype_id,
				b.card_id,
				COUNT(b.CardAmount) AS amount,
				SUM(b.CardAmount) AS deck_amount
			FROM
				(
					SELECT
						da.archetype_id,
						da.deck_id,
						c_d.card_id,
						COUNT(c_d.card_id) AS CardAmount
					FROM
						mv_deck_archetypes da
						LEFT JOIN graph_cards_belong_to_decks c_d ON c_d.deck_id = da.deck_id
					WHERE
						da.archetype_id = archetype_id
						AND da.weight > 5
					GROUP BY
						da.deck_id,
						c_d.card_id,
						da.archetype_id
				) b
			JOIN entity_cards ec ON ec.id = b.card_id
			WHERE NOT match_archetype(ec.set_code, archetype_id)
			GROUP BY b.card_id, b.archetype_id
			LIMIT 100;
		END;
	`)
}
