package config

import (
	"fmt"
	"os"

	model "github.com/lplanch/test-go-api/models"
	util "github.com/lplanch/test-go-api/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connection() *gorm.DB {
	databaseURI := make(chan string, 1)
	config := gorm.Config{}

	if os.Getenv("RAILWAY_ENVIRONMENT_NAME") != "production" {
		databaseURI <- fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			util.GodotEnv("DB_USER"), util.GodotEnv("DB_PASSWORD"), util.GodotEnv("DB_HOST"), util.GodotEnv("DB_PORT"), util.GodotEnv("DB_NAME"))
		config.Logger = logger.Default.LogMode(logger.Info)
	} else {
		databaseURI <- fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	}

	db, err := gorm.Open(mysql.Open(<-databaseURI), &config)

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("RAILWAY_ENVIRONMENT_NAME") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		// KV
		&model.KeyValueStore{},
		// ENTITIES
		&model.EntityCard{},
		&model.EntityDeck{},
		&model.EntityBanlist{},
		// MATERIALIZED VIEWS
		&model.MvTopCard{},
		&model.MvDeckArchetypes{},
		&model.MvTopArchetype{},
		&model.MvTopArchetypeCard{},
		&model.MvTopRelatedCards{},
		// GRAPHS
		&model.GraphCardsBelongToDecks{},
		&model.GraphCardsBelongToBanlists{},
		// ENUMS
		&model.EnumRule{},
		&model.EnumAttribute{},
		&model.EnumLevel{},
		&model.EnumCategory{},
		&model.EnumRace{},
		&model.EnumType{},
		&model.EnumArchetype{},
	)

	// FUNCTIONS
	model.AutoMigrateFunctionMatchArchetype(db)

	// PROCEDURES
	model.AutoMigrateProcedureMvTopCards(db)
	model.AutoMigrateProcedureMvDeckArchetypes(db)
	model.AutoMigrateProcedureMvTopArchetypes(db)
	model.AutoMigrateProcedureMvTopArchetypeCards(db)

	// TRIGGERS
	model.AutoMigrateTriggerMvDeckArchetypes(db)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
