package getVersion

import (
	model "github.com/lplanch/test-go-api/models"
	util "github.com/lplanch/test-go-api/utils"
	"gorm.io/gorm"
)

type Repository interface {
	GetVersionRepository() (*[]model.KeyValueStore, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryGet(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetVersionRepository() (*[]model.KeyValueStore, string) {

	var kv []model.KeyValueStore
	db := r.db.Model(&kv)
	errorCode := make(chan string, 1)

	getKv := db.Debug().Where("`key` = ?", util.KV_ENUMS_LAST_COMMIT).Or("`key` = ?", util.KV_BABELCDB_LAST_COMMIT).Or("`key` = ?", util.KV_ENUMS_VERSION_DATE).Or("`key` = ?", util.KV_BABELCDB_VERSION_DATE).Find(&kv)

	if getKv.RowsAffected < 1 {
		errorCode <- "VERSION_NOT_FOUND_500"
		return &kv, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &kv, <-errorCode
}
