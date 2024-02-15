package route

import (
	"github.com/gin-gonic/gin"
	getHealthcheck "github.com/lplanch/test-go-api/controllers/misc-controllers/healthcheck"
	handlerGetHealthcheck "github.com/lplanch/test-go-api/handlers/misc-handlers/healthcheck"
	"gorm.io/gorm"
)

func InitMiscRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Misc
	*/
	getHealthcheckService := getHealthcheck.NewServiceGet()
	healthcheckHandler := handlerGetHealthcheck.NewHandlerGetHealthcheck(getHealthcheckService)

	/**
	@description All Misc Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/health", healthcheckHandler.HealthcheckHandler)
}
