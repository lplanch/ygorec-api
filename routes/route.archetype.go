package route

import (
	"github.com/gin-gonic/gin"
	listArchetypes "github.com/lplanch/test-go-api/controllers/archetype-controllers/list"
	handlerListArchetypes "github.com/lplanch/test-go-api/handlers/archetype-handlers/list"
	"gorm.io/gorm"
)

func InitArchetypeRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Archetype
	*/

	listArchetypesRepository := listArchetypes.NewRepositoryList(db)
	listArchetypesService := listArchetypes.NewServiceList(listArchetypesRepository)
	listArchetypesHandler := handlerListArchetypes.NewHandlerListArchetypes(listArchetypesService)

	/**
	@description All Archetype Routes
	*/
	groupRoute := route.Group("/api")
	groupRoute.GET("/archetypes", listArchetypesHandler.ListArchetypesHandler)
}
