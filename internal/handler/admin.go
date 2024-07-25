package handler

import (
	"github.com/gin-gonic/gin"
	"novelman/internal/service"
)

type AdminHandler struct {
	*Handler
	adminService service.AdminService
}

func NewAdminHandler(
    handler *Handler,
    adminService service.AdminService,
) *AdminHandler {
	return &AdminHandler{
		Handler:      handler,
		adminService: adminService,
	}
}

func (h *AdminHandler) GetAdmin(ctx *gin.Context) {

}
