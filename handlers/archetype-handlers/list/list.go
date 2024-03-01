package handlerListArchetypes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	listArchetypes "github.com/lplanch/test-go-api/controllers/archetype-controllers/list"
	util "github.com/lplanch/test-go-api/utils"
)

type handler struct {
	service listArchetypes.Service
}

func NewHandlerListArchetypes(service listArchetypes.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ListArchetypesHandler(ctx *gin.Context) {
	input := listArchetypes.InputListArchetypes{Limit: 20, Offset: 0}

	if len(ctx.Query("limit")) > 0 {
		input.Limit, _ = strconv.Atoi(ctx.Query("limit"))
	}
	if len(ctx.Query("offset")) > 0 {
		input.Offset, _ = strconv.Atoi(ctx.Query("offset"))
	}

	err := util.Validate(&input)

	if err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, err)
		return
	}

	listArchetypes := h.service.ListArchetypesService(&input)

	util.APIResponse(ctx, "Archetypes list found", http.StatusOK, http.MethodGet, listArchetypes)
}
