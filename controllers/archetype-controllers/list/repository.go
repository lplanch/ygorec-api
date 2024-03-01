package listArchetypes

import (
	model "github.com/lplanch/test-go-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	ListArchetypesRepository(input *InputListArchetypes) *[]model.ModelArchetype
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryList(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ListArchetypesRepository(input *InputListArchetypes) *[]model.ModelArchetype {

	var archetypes []model.ModelArchetype

	db := r.db.Model(&model.MvDeckArchetypes{})

	db.Debug().Select(`
		mv_deck_archetypes.archetype_id,
		a.value AS label,
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
	`).Where(`
		mv_deck_archetypes.weight > 5
	`).Joins(`
		JOIN enum_archetypes a ON a.id = mv_deck_archetypes.archetype_id
	`).Group("archetype_id").Order("deck_amount DESC, card_amount DESC").Limit(input.Limit).Offset(input.Offset).Find(&archetypes)

	return &archetypes
}
