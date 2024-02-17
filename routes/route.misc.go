package route

import (
	"github.com/gin-gonic/gin"
	getHealthcheck "github.com/lplanch/test-go-api/controllers/misc-controllers/healthcheck"
	getVersion "github.com/lplanch/test-go-api/controllers/misc-controllers/version"
	handlerGetHealthcheck "github.com/lplanch/test-go-api/handlers/misc-handlers/healthcheck"
	handlerGetVersion "github.com/lplanch/test-go-api/handlers/misc-handlers/version"
	"gorm.io/gorm"
)

func InitMiscRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Misc
	*/
	getHealthcheckService := getHealthcheck.NewServiceGet()
	healthcheckHandler := handlerGetHealthcheck.NewHandlerGetHealthcheck(getHealthcheckService)

	getVersionRepository := getVersion.NewRepositoryGet(db)
	getVersionService := getVersion.NewServiceGet(getVersionRepository)
	getVersionHandler := handlerGetVersion.NewHandlerGetVersion(getVersionService)

	/**
	@description All Misc Route
	*/
	groupRoute := route.Group("/api")
	groupRoute.GET("/health", healthcheckHandler.HealthcheckHandler)
	groupRoute.GET("/version", getVersionHandler.GetVersionHandler)
}
