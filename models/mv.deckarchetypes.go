package model

import "gorm.io/gorm"

type MvDeckArchetypes struct {
	DeckID      string `gorm:"index:idx_mv_deck_archetypes,unique"`
	ArchetypeID uint64 `gorm:"index:idx_mv_deck_archetypes,unique"`
	Weight      uint8  `gorm:"type:TINYINT UNSIGNED;NOT NULL"`

	// Constraint
	Deck      EntityDeck    `gorm:"foreignKey:DeckID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Archetype EnumArchetype `gorm:"foreignKey:ArchetypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func AutoMigrateProcedureMvDeckArchetypes(db *gorm.DB) {
	db.Exec(`
		DROP PROCEDURE IF EXISTS refresh_mv_deck_archetype;
	`)
	db.Exec(`
		CREATE PROCEDURE refresh_mv_deck_archetype (IN deck_id VARCHAR(255))
		BEGIN
			DELETE FROM mv_deck_archetypes mv WHERE mv.deck_id = deck_id;
			INSERT INTO mv_deck_archetypes
			SELECT
				cd.deck_id,
				a.id AS archetype_id,
				COUNT(c.id) AS weight
			FROM
				enum_archetypes a
				JOIN graph_cards_belong_to_decks cd ON cd.deck_id = deck_id
				JOIN entity_cards c
					ON c.id = cd.card_id AND
					(
						(c.set_code & 65535) = a.id OR (((c.set_code & 65535) & 4095) = (a.id & 4095) AND ((c.set_code & 65535) & a.id) = (a.id & 65535)) OR
						((c.set_code >> 16) & 65535) = a.id OR ((((c.set_code >> 16) & 65535) & 4095) = (a.id & 4095) AND (((c.set_code >> 16) & 65535) & a.id) = (a.id & 65535)) OR
						((c.set_code >> 32) & 65535) = a.id OR ((((c.set_code >> 32) & 65535) & 4095) = (a.id & 4095) AND (((c.set_code >> 32) & 65535) & a.id) = (a.id & 65535)) OR
						(c.set_code >> 48) = a.id OR (((c.set_code >> 48) & 4095) = (a.id & 4095) AND ((c.set_code >> 48) & a.id) = (a.id & 65535))
					)
				GROUP BY a.id, cd.deck_id;
		END;
	`)
}

func AutoMigrateTriggerMvDeckArchetypes(db *gorm.DB) {
	db.Exec(`
		DROP TRIGGER IF EXISTS trigger_mv_deck_archetype_insert;
	`)
	db.Exec(`
		DROP TRIGGER IF EXISTS trigger_mv_deck_archetype_update;
	`)
	db.Exec(`
		CREATE TRIGGER trigger_mv_deck_archetype_insert
			AFTER INSERT ON entity_decks
			FOR EACH ROW
		BEGIN
			CALL refresh_mv_deck_archetype(NEW.id);
		END;
	`)
	db.Exec(`
		CREATE TRIGGER trigger_mv_deck_archetype_update
			AFTER UPDATE ON entity_decks
			FOR EACH ROW
		BEGIN
			CALL refresh_mv_deck_archetype(NEW.id);
		END;	
	`)
}
