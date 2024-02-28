package handlerSearchCard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	searchCard "github.com/lplanch/test-go-api/controllers/card-controllers/search"
	util "github.com/lplanch/test-go-api/utils"
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

	err := util.Validate(&input)

	if err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, err)
		return
	}

	getCard := h.service.SearchCardService(&input)

	util.APIResponse(ctx, "Search cards data found", http.StatusOK, http.MethodGet, getCard)
}
