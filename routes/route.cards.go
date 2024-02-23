package route

import (
	"github.com/gin-gonic/gin"
	getCard "github.com/lplanch/test-go-api/controllers/card-controllers/get"
	listCards "github.com/lplanch/test-go-api/controllers/card-controllers/list"
	searchCard "github.com/lplanch/test-go-api/controllers/card-controllers/search"
	handlerGetCard "github.com/lplanch/test-go-api/handlers/card-handlers/get"
	handlerListCards "github.com/lplanch/test-go-api/handlers/card-handlers/list"
	handlerSearchCard "github.com/lplanch/test-go-api/handlers/card-handlers/search"
	"gorm.io/gorm"
)

func InitCardsRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Cards
	*/

	searchCardRepository := searchCard.NewRepositoryList(db)
	searchCardService := searchCard.NewServiceList(searchCardRepository)
	searchCardHandler := handlerSearchCard.NewHandlerSearchCard(searchCardService)

	getCardRepository := getCard.NewRepositoryGet(db)
	getCardService := getCard.NewServiceGet(getCardRepository)
	getCardHandler := handlerGetCard.NewHandlerGetCard(getCardService)

	listCardsRepository := listCards.NewRepositoryList(db)
	listCardsService := listCards.NewServiceList(listCardsRepository)
	listCardsHandler := handlerListCards.NewHandlerListCards(listCardsService)

	/**
	@description All Cards Route
	*/
	groupRoute := route.Group("/api")
	groupRoute.GET("/typeahead", searchCardHandler.SearchCardHandler)
	groupRoute.GET("/cards/:id", getCardHandler.GetCardHandler)
	groupRoute.GET("/cards", listCardsHandler.ListCardsHandler)
}
