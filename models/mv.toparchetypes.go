package model

import "gorm.io/gorm"

type MvTopArchetype struct {
	ArchetypeID    uint64 `gorm:"index:idx_top_archetype_archetype,unique"`
	DeckAmount     uint32 `gorm:"index:INT UNSIGNED;NOT NULL;DEFAULT NULL"`
	CardAmount     uint32 `gorm:"index:INT UNSIGNED;NOT NULL;DEFAULT NULL"`
	MostUsedCardID uint64 `gorm:"index:idx_top_archetype_card"`

	// Constraint
	Archetype EnumArchetype `gorm:"foreignKey:ArchetypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Card      EntityCard    `gorm:"foreignKey:MostUsedCardID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func AutoMigrateProcedureMvTopArchetypes(db *gorm.DB) {
	db.Exec(`
		DROP PROCEDURE IF EXISTS refresh_mv_top_archetypes;
	`)
	db.Exec(`
		CREATE PROCEDURE refresh_mv_top_archetypes ()
		BEGIN
			DELETE FROM mv_top_archetypes mv;
			INSERT INTO mv_top_archetypes
			SELECT 
				mv_deck_archetypes.archetype_id,
				COUNT(mv_deck_archetypes.archetype_id) AS deck_amount,
				SUM(mv_deck_archetypes.weight) AS card_amount,
				(
					SELECT card_id
					FROM mv_top_cards
					LEFT OUTER JOIN entity_cards e ON e.id = card_id
					WHERE match_archetype(e.set_code, mv_deck_archetypes.archetype_id)
					ORDER BY amount DESC
					LIMIT 1
				) AS most_used_card_id
			FROM mv_deck_archetypes
			WHERE 
				mv_deck_archetypes.weight > 5
			GROUP BY archetype_id;
		END;
	`)
}
