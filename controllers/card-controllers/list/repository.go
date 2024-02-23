package listCards

import (
	model "github.com/lplanch/test-go-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	ListCardsRepository(input *InputListCards) *[]model.ModelListCardStats
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryList(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ListCardsRepository(input *InputListCards) *[]model.ModelListCardStats {

	var cards []model.ModelListCardStats

	db := r.db.Model(&model.MvTopCard{})

	db.Debug().Select(`
		mv_top_cards.card_id AS id,
		e.name AS label,
		CONCAT('/cards/', CONVERT(mv_top_cards.card_id, char)) AS url,
		mv_top_cards.percentage,
		mv_top_cards.average
	`).Joins(`
		JOIN entity_cards e ON e.id = mv_top_cards.card_id
	`).Order(`
		mv_top_cards.percentage DESC,
		mv_top_cards.card_id ASC
	`).Limit(input.Limit).Offset(input.Offset).Find(&cards)

	return &cards
}
