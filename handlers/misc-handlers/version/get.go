package handlerGetVersion

import (
	"net/http"

	"github.com/gin-gonic/gin"
	getVersion "github.com/lplanch/test-go-api/controllers/misc-controllers/version"
	util "github.com/lplanch/test-go-api/utils"
)

type handler struct {
	service getVersion.Service
}

func NewHandlerGetVersion(service getVersion.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResultStudentHandler(ctx *gin.Context) {

	resultStudent, errResultStudent := h.service.GetVersionService()

	switch errResultStudent {

	default:
		util.APIResponse(ctx, "Get version data success", http.StatusOK, http.MethodGet, resultStudent)
	}
}
