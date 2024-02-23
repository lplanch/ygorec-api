package handlerListBanlists

import (
	"net/http"

	"github.com/gin-gonic/gin"
	listBanlists "github.com/lplanch/test-go-api/controllers/banlist-controllers/list"
	util "github.com/lplanch/test-go-api/utils"
)

type handler struct {
	service listBanlists.Service
}

func NewHandlerListBanlists(service listBanlists.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ListBanlistsHandler(ctx *gin.Context) {

	listBanlists := h.service.ListBanlistsService()

	util.APIResponse(ctx, "Banlists data found", http.StatusOK, http.MethodGet, listBanlists)
}
