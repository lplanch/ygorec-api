package handlerGetHealthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
	getHealthcheck "github.com/lplanch/test-go-api/controllers/misc-controllers/healthcheck"
	util "github.com/lplanch/test-go-api/utils"
)

type handler struct {
	service getHealthcheck.Service
}

func NewHandlerGetHealthcheck(service getHealthcheck.Service) *handler {
	return &handler{service: service}
}

func (h *handler) HealthcheckHandler(ctx *gin.Context) {

	resultGetHealthcheck := h.service.GetHealthcheckService()

	util.APIResponse(ctx, "Healthcheck successful", http.StatusOK, http.MethodGet, resultGetHealthcheck)
}
