package handler

import (
	"novelman/internal/handler"
	"novelman/internal/service"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	*handler.Handler
	adminService service.AdminService
}

func NewAdminHandler(
	handler *handler.Handler,
	adminService service.AdminService,
) *AdminHandler {
	return &AdminHandler{
		Handler:      handler,
		adminService: adminService,
	}
}

func (h *AdminHandler) GetAdmin(ctx *gin.Context) {

}
