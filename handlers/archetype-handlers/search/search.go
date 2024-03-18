package handlerSearchArchetype

import (
	"net/http"

	"github.com/gin-gonic/gin"
	searchArchetype "github.com/lplanch/test-go-api/controllers/archetype-controllers/search"
	util "github.com/lplanch/test-go-api/utils"
)

type handler struct {
	service searchArchetype.Service
}

func NewHandlerSearchArchetype(service searchArchetype.Service) *handler {
	return &handler{service: service}
}

func (h *handler) SearchArchetypeHandler(ctx *gin.Context) {

	var input searchArchetype.InputSearchArchetype
	input.Q = ctx.Query("q")

	err := util.Validate(&input)

	if err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, err)
		return
	}

	searchArchetypes := h.service.SearchArchetypeService(&input)

	util.APIResponse(ctx, "Search archetype data found", http.StatusOK, http.MethodGet, searchArchetypes)
}
