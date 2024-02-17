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

func (h *handler) GetVersionHandler(ctx *gin.Context) {

	getVersion, errGetVersion := h.service.GetVersionService()

	switch errGetVersion {

	case "VERSION_NOT_FOUND_500":
		util.APIResponse(ctx, "Version not found", http.StatusInternalServerError, http.MethodGet, nil)
	default:
		util.APIResponse(ctx, "Get version data success", http.StatusOK, http.MethodGet, getVersion)
	}
}
