package getVersion

import (
	"time"

	model "github.com/lplanch/test-go-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetVersionRepository() (*model.StaticVersion, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryGet(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetVersionRepository() (*model.StaticVersion, string) {

	var version model.StaticVersion = model.StaticVersion{LastCommit: "last_commit_sha", LastUpdate: time.Now()}
	// db := r.db
	errorCode := make(chan string, 1)

	// getVersion := db.Debug().Select("*")

	// if getVersion.Error != nil {
	// 	errorCode <- "RESULTS_STUDENT_NOT_FOUND_404"
	// 	return &version, <-errorCode
	// } else {
	errorCode <- "nil"
	// }

	return &version, <-errorCode
}
