package handlerGetArchetype

import (
	"net/http"

	"github.com/gin-gonic/gin"
	getArchetype "github.com/lplanch/test-go-api/controllers/archetype-controllers/get"
	util "github.com/lplanch/test-go-api/utils"
)

type handler struct {
	service getArchetype.Service
}

func NewHandlerGetArchetype(service getArchetype.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetArchetypeHandler(ctx *gin.Context) {

	var input = getArchetype.InputGetArchetype{}
	input.Value = ctx.Param("value")

	err := util.Validate(&input)

	if err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, err)
		return
	}

	input_service, errGetArchetypeID := h.service.GetArchetypeIDFromNameService(&input)

	switch errGetArchetypeID {

	case "GET_ARCHETYPE_INPUT_SERVICE_NOT_FOUND_404":
		util.APIResponse(ctx, "Archetype not found", http.StatusNotFound, http.MethodGet, nil)
		return
	}

	getArchetype := h.service.GetArchetypeService(input_service)

	util.APIResponse(ctx, "Archetype data found", http.StatusOK, http.MethodGet, getArchetype)
}
