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

	db := r.db.Model(&model.MvTopArchetype{})

	db.Debug().Select(`
		archetype_id,
		a.value AS label,
		deck_amount,
		card_amount,
		most_used_card_id
	`).Joins(`
		JOIN enum_archetypes a ON a.id = archetype_id
	`).Order("deck_amount DESC, card_amount DESC, label ASC").Limit(input.Limit).Offset(input.Offset).Find(&archetypes)

	return &archetypes
}
