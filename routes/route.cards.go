package route

import (
	"github.com/gin-gonic/gin"
	getCard "github.com/lplanch/test-go-api/controllers/card-controllers/get"
	handlerGetCard "github.com/lplanch/test-go-api/handlers/card-handlers/get"
	"gorm.io/gorm"
)

func InitCardsRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Misc
	*/

	getCardRepository := getCard.NewRepositoryGet(db)
	getCardService := getCard.NewServiceGet(getCardRepository)
	getCardHandler := handlerGetCard.NewHandlerGetCard(getCardService)

	/**
	@description All Misc Route
	*/
	groupRoute := route.Group("/api")
	groupRoute.GET("/cards/:id", getCardHandler.GetCardHandler)
}
