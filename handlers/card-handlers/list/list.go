package handlerListCards

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	listCards "github.com/lplanch/test-go-api/controllers/card-controllers/list"
	util "github.com/lplanch/test-go-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handler struct {
	service listCards.Service
}

func NewHandlerListCards(service listCards.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ListCardsHandler(ctx *gin.Context) {

	var input listCards.InputListCards
	input.From, _ = time.Parse(time.RFC3339, ctx.Query("from"))
	input.To, _ = time.Parse(time.RFC3339, ctx.Query("to"))

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			{
				Tag:     "datetime",
				Field:   "From",
				Message: "from must be a date if specified",
			},
			{
				Tag:     "datetime",
				Field:   "To",
				Message: "to must be a date if specified",
			},
			{
				Tag:     "gt",
				Field:   "Limit",
				Message: "limit must be > 0 if specified",
			},
			{
				Tag:     "gte",
				Field:   "Offset",
				Message: "offset must be >= 0 if specified",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	getCard, _ := h.service.ListCardsService(&input)

	util.APIResponse(ctx, "List cards data found", http.StatusOK, http.MethodGet, getCard)
}
