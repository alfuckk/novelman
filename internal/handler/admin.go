package handler

import (
	"net/http"
	v1 "novelman/api/v1"
	"novelman/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

}

// Register godoc
// @Summary Register Admin
// @Schemes
// @Description 目前只支持邮箱登录
// @Tags Admin
// @Accept json
// @Produce json
// @Param request body v1.RegisterRequest true "params"
// @Success 200 {object} v1.Response
// @Router /register [post]
func (h *AdminHandler) Register(ctx *gin.Context) {
	req := new(v1.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if err := h.adminService.Register(ctx, req); err != nil {
		h.logger.WithContext(ctx).Error("adminService.Register error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

func (h *AdminHandler) GetAdmin(ctx *gin.Context) {

}
