package listCards

import (
	model "github.com/lplanch/test-go-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	ListCardsRepository(input *InputListCards) (*[]model.EntityCard, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryGet(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ListCardsRepository(input *InputListCards) (*[]model.EntityCard, string) {

	var cards []model.EntityCard

	db := r.db.Model(&model.EntityCard{})
	errorCode := make(chan string, 1)

	listCards := db.Debug().Limit(20).Find(&cards)

	if listCards.RowsAffected < 1 {
		errorCode <- "GET_CARD_NOT_FOUND_404"
		return &cards, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &cards, <-errorCode
}
