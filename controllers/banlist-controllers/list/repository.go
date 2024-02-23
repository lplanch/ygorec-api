package listBanlists

import (
	model "github.com/lplanch/test-go-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	ListBanlistsRepository() *[]model.ModelBanlist
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryList(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ListBanlistsRepository() *[]model.ModelBanlist {

	var banlists []model.ModelBanlist

	db := r.db.Model(&model.EntityBanlist{})

	db.Debug().Select(`
		STR_TO_DATE(entity_banlists.id, '%Y-%m-%d') AS date,
		entity_banlists.ot
	`).Find(&banlists)

	return &banlists
}
