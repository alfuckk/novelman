package handler

import (
	"net/http"
	v1 "novelman/api/v1"
	"novelman/internal/service"

	"github.com/gin-gonic/gin"
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

// Login godoc
// @Summary Login Admin
// @Schemes
// @Description
// @Tags Admin
// @Accept json
// @Produce json
// @Param request body v1.LoginRequest true "params"
// @Success 200 {object} v1.LoginResponse
// @Router /login [post]
func (h *AdminHandler) Login(ctx *gin.Context) {
	var req v1.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	token, err := h.adminService.Login(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}
	v1.HandleSuccess(ctx, v1.LoginResponseData{
		AccessToken: token,
	})
}

func (h *AdminHandler) GetAdmin(ctx *gin.Context) {

}
