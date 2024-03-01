package main

import (
	"log"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	config "github.com/lplanch/test-go-api/configs"
	route "github.com/lplanch/test-go-api/routes"
	util "github.com/lplanch/test-go-api/utils"
)

func main() {
	/**
	@description Utils initialisation
	*/
	util.InitUptime()

	/**
	@description Setup Server
	*/
	router := SetupRouter()

	/**
	@description Run Server
	*/
	log.Fatal(router.Run(":" + util.GodotEnv("PORT")))
}

func SetupRouter() *gin.Engine {
	/**
	@description Setup Database Connection
	*/
	db := config.Connection()

	/**
	@description Init Router
	*/
	router := gin.Default()

	/**
	@description Setup Mode Application
	*/
	env_name := util.GodotEnv("RAILWAY_ENVIRONMENT_NAME")
	println(env_name)
	if env_name == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else if env_name == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	/**
	@description Validator Initialization
	*/
	util.InitValidator()

	/**
	@description Setup Middleware
	*/
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))

	/**
	@description Init All Route
	*/
	route.InitMiscRoutes(db, router)
	route.InitBanlistRoutes(db, router)
	route.InitArchetypeRoutes(db, router)
	route.InitCardsRoutes(db, router)

	return router
}
