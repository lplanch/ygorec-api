package searchCard

import (
	model "github.com/lplanch/test-go-api/models"
	util "github.com/lplanch/test-go-api/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	SearchCardRepository(input *InputSearchCard) *[]model.ModelListCard
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryList(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) SearchCardRepository(input *InputSearchCard) *[]model.ModelListCard {

	var cards []model.ModelListCard

	db := r.db.Model(&model.EntityCard{})

	db.Debug().Select(`
		id,
		name AS label,
		CONCAT('/cards/', CONVERT(id, char)) AS url
	`).Where("name LIKE ?", "%"+input.Q+"%").Clauses(
		util.OrderByCase{
			Column: clause.Column{Name: "name"},
			Values: map[string]int{
				input.Q:             0,
				input.Q + "%":       1,
				"%" + input.Q:       2,
				"%" + input.Q + "%": 3,
			},
			Asc: true,
		},
	).Limit(10).Find(&cards)

	return &cards
}
