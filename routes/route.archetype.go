package route

import (
	"github.com/gin-gonic/gin"
	getArchetype "github.com/lplanch/test-go-api/controllers/archetype-controllers/get"
	listArchetypes "github.com/lplanch/test-go-api/controllers/archetype-controllers/list"
	searchArchetype "github.com/lplanch/test-go-api/controllers/archetype-controllers/search"
	handlerGetArchetype "github.com/lplanch/test-go-api/handlers/archetype-handlers/get"
	handlerListArchetypes "github.com/lplanch/test-go-api/handlers/archetype-handlers/list"
	handlerSearchArchetype "github.com/lplanch/test-go-api/handlers/archetype-handlers/search"
	"gorm.io/gorm"
)

func InitArchetypeRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Archetype
	*/

	searchArchetypeRepository := searchArchetype.NewRepositoryList(db)
	searchArchetypeService := searchArchetype.NewServiceList(searchArchetypeRepository)
	searchArchetypeHandler := handlerSearchArchetype.NewHandlerSearchArchetype(searchArchetypeService)

	listArchetypesRepository := listArchetypes.NewRepositoryList(db)
	listArchetypesService := listArchetypes.NewServiceList(listArchetypesRepository)
	listArchetypesHandler := handlerListArchetypes.NewHandlerListArchetypes(listArchetypesService)

	getArchetypeRepository := getArchetype.NewRepositoryGet(db)
	getArchetypeService := getArchetype.NewServiceGet(getArchetypeRepository)
	getArchetypeHandler := handlerGetArchetype.NewHandlerGetArchetype(getArchetypeService)

	/**
	@description All Archetype Routes
	*/
	groupRoute := route.Group("/api")
	groupRoute.GET("/archetypes/typeahead", searchArchetypeHandler.SearchArchetypeHandler)
	groupRoute.GET("/archetypes", listArchetypesHandler.ListArchetypesHandler)
	groupRoute.GET("/archetypes/:value", getArchetypeHandler.GetArchetypeHandler)
}
