package route

import (
	"github.com/gin-gonic/gin"
	listBanlists "github.com/lplanch/test-go-api/controllers/banlist-controllers/list"
	handlerListBanlists "github.com/lplanch/test-go-api/handlers/banlist-handlers/list"
	"gorm.io/gorm"
)

func InitBanlistRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Banlist
	*/

	listBanlistsRepository := listBanlists.NewRepositoryList(db)
	listBanlistsService := listBanlists.NewServiceList(listBanlistsRepository)
	listBanlistsHandler := handlerListBanlists.NewHandlerListBanlists(listBanlistsService)

	/**
	@description All Banlist Routes
	*/
	groupRoute := route.Group("/api")
	groupRoute.GET("/banlists", listBanlistsHandler.ListBanlistsHandler)
}
