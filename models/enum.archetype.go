package model

import "gorm.io/gorm"

type EnumArchetype struct {
	ID    uint64 `gorm:"type:BIGINT UNSIGNED PRIMARY KEY"`
	Value string `gorm:"type:TEXT;NOT NULL"`
}

func AutoMigrateFunctionMatchArchetype(db *gorm.DB) {
	db.Exec(`
		DROP FUNCTION IF EXISTS match_archetype;
	`)
	db.Exec(`
		CREATE FUNCTION match_archetype (
			set_code BIGINT UNSIGNED,
			archetype_id BIGINT UNSIGNED
		)
		RETURNS BOOLEAN
		NO SQL
		BEGIN
			DECLARE flag BOOLEAN;
			SET flag = ((set_code & 65535) = archetype_id OR (((set_code & 65535) & 4095) = (archetype_id & 4095) AND ((set_code & 65535) & archetype_id) = (archetype_id & 65535)) OR
			((set_code >> 16) & 65535) = archetype_id OR ((((set_code >> 16) & 65535) & 4095) = (archetype_id & 4095) AND (((set_code >> 16) & 65535) & archetype_id) = (archetype_id & 65535)) OR
			((set_code >> 32) & 65535) = archetype_id OR ((((set_code >> 32) & 65535) & 4095) = (archetype_id & 4095) AND (((set_code >> 32) & 65535) & archetype_id) = (archetype_id & 65535)) OR
			(set_code >> 48) = archetype_id OR (((set_code >> 48) & 4095) = (archetype_id & 4095) AND ((set_code >> 48) & archetype_id) = (archetype_id & 65535)));
			RETURN flag;
		END;
	`)
}
