package config

import (
	"os"

	model "github.com/lplanch/test-go-api/models"
	util "github.com/lplanch/test-go-api/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connection() *gorm.DB {
	databaseURI := make(chan string, 1)
	config := gorm.Config{}

	if os.Getenv("RAILWAY_ENVIRONMENT_NAME") != "production" {
		databaseURI <- util.GodotEnv("DATABASE_PATH")
		config.Logger = logger.Default.LogMode(logger.Info)
	} else {
		databaseURI <- os.Getenv("DATABASE_PATH")
	}

	db, err := gorm.Open(sqlite.Open(<-databaseURI), &config)

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("RAILWAY_ENVIRONMENT_NAME") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		&model.KeyValueStore{},
		&model.EntityCard{},
		&model.EnumRule{},
		&model.EnumAttribute{},
		&model.EnumLevel{},
		&model.EnumCategory{},
		&model.EnumRace{},
		&model.EnumType{},
		&model.EnumArchetype{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
