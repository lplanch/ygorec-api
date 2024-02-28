package handlerGetCard

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	getCard "github.com/lplanch/test-go-api/controllers/card-controllers/get"
	util "github.com/lplanch/test-go-api/utils"
)

type handler struct {
	service getCard.Service
}

func NewHandlerGetCard(service getCard.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetCardHandler(ctx *gin.Context) {

	var input getCard.InputGetCard
	input.ID = ctx.Param("id")

	err := util.Validate(&input)

	if err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, err)
		return
	}

	uint_id, _ := strconv.ParseUint(input.ID, 10, 64)

	var input_service = getCard.InputServiceGetCard{ID: uint_id}

	getCard, errGetCard := h.service.GetCardService(&input_service)

	switch errGetCard {

	case "GET_CARD_NOT_FOUND_404":
		util.APIResponse(ctx, "Card data not found", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		util.APIResponse(ctx, "Card data found", http.StatusOK, http.MethodGet, getCard)
	}
}
