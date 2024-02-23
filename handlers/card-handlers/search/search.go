package handlerSearchCard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	searchCard "github.com/lplanch/test-go-api/controllers/card-controllers/search"
	util "github.com/lplanch/test-go-api/utils"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handler struct {
	service searchCard.Service
}

func NewHandlerSearchCard(service searchCard.Service) *handler {
	return &handler{service: service}
}

func (h *handler) SearchCardHandler(ctx *gin.Context) {

	var input searchCard.InputSearchCard
	input.Q = ctx.Query("q")

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "Q",
				Message: "q is required in query",
			},
			{
				Tag:     "lowercase",
				Field:   "Q",
				Message: "q must be a lowercase string",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	getCard := h.service.SearchCardService(&input)

	util.APIResponse(ctx, "Search cards data found", http.StatusOK, http.MethodGet, getCard)
}
