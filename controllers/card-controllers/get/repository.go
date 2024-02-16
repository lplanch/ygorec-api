package getCard

import (
	model "github.com/lplanch/test-go-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetCardRepository(input *model.EntityCard) (*model.EntityCard, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryGet(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetCardRepository(input *model.EntityCard) (*model.EntityCard, string) {

	var card model.EntityCard

	db := r.db.Model(&card)
	errorCode := make(chan string, 1)

	getCard := db.Debug().Select("*").Where("id = ?", input.ID).Find(&card)

	if getCard.RowsAffected < 1 {
		errorCode <- "GET_CARD_NOT_FOUND_404"
		return &card, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &card, <-errorCode
}
