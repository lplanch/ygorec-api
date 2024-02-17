package handlerGetCard

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	getCard "github.com/lplanch/test-go-api/controllers/card-controllers/get"
	util "github.com/lplanch/test-go-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
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

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "ID",
				Message: "id is required in path",
			},
			{
				Tag:     "number",
				Field:   "ID",
				Message: "id must be a number",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)
	uint_id, _ := strconv.ParseUint(input.ID, 10, 64)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

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
