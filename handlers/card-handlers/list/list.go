package handlerListCards

import (
	"net/http"
	"strconv"

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

	input := listCards.InputListCards{Limit: 20, Offset: 0, Banlist: ""}

	if len(ctx.Query("limit")) > 0 {
		input.Limit, _ = strconv.Atoi(ctx.Query("limit"))
	}
	if len(ctx.Query("offset")) > 0 {
		input.Offset, _ = strconv.Atoi(ctx.Query("offset"))
	}
	if len(ctx.Query("banlist")) > 0 {
		input.Banlist = ctx.Query("banlist")
	}

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
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
			{
				Tag:     "is-awesome",
				Field:   "Banlist",
				Message: "banlist must be a valid banlist",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	getCard := h.service.ListCardsService(&input)

	util.APIResponse(ctx, "List cards data found", http.StatusOK, http.MethodGet, getCard)
}
