package searchArchetype

import (
	model "github.com/lplanch/test-go-api/models"
	util "github.com/lplanch/test-go-api/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	SearchArchetypeRepository(input *InputSearchArchetype) *[]model.ModelListArchetype
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryList(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) SearchArchetypeRepository(input *InputSearchArchetype) *[]model.ModelListArchetype {

	var archetypes []model.ModelListArchetype

	db := r.db.Model(&model.EnumArchetype{})

	db.Debug().Select(`
		id,
		value AS label,
		CONCAT('/archetypes/', CONVERT(LOWER(REPLACE(value, ' ', '-')), char)) AS url
	`).Where("value LIKE ?", "%"+input.Q+"%").Clauses(
		util.OrderByCase{
			Column: clause.Column{Name: "value"},
			Values: map[string]int{
				input.Q:             0,
				input.Q + "%":       1,
				"%" + input.Q:       2,
				"%" + input.Q + "%": 3,
			},
			Asc: true,
		},
	).Limit(10).Find(&archetypes)

	return &archetypes
}
