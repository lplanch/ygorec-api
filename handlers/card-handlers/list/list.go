package handlerListCards

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	listCards "github.com/lplanch/test-go-api/controllers/card-controllers/list"
	util "github.com/lplanch/test-go-api/utils"
)

type handler struct {
	service listCards.Service
}

func NewHandlerListCards(service listCards.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ListCardsHandler(ctx *gin.Context) {

	input := listCards.InputListCards{Limit: 20, Offset: 0, CardID: 0}

	if len(ctx.Query("limit")) > 0 {
		input.Limit, _ = strconv.Atoi(ctx.Query("limit"))
	}
	if len(ctx.Query("offset")) > 0 {
		input.Offset, _ = strconv.Atoi(ctx.Query("offset"))
	}
	if len(ctx.Query("banlist")) > 0 {
		input.Banlist = ctx.Query("banlist")
	}
	if len(ctx.Query("card_id")) > 0 {
		input.CardID, _ = strconv.Atoi(ctx.Query("card_id"))
	}

	err := util.Validate(&input)

	if err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, err)
		return
	}

	getCard := h.service.ListCardsService(&input)

	util.APIResponse(ctx, "List cards data found", http.StatusOK, http.MethodGet, getCard)
}
