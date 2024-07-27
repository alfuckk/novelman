package handler

import (
	"net/http"
	v1 "novelman/api/v1"
	"novelman/internal/service"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	*Handler
	appService service.AppService
}

func NewAppHandler(
	handler *Handler,
	appService service.AppService,
) *AppHandler {
	return &AppHandler{
		Handler:    handler,
		appService: appService,
	}
}

func (h *AppHandler) GetApp(ctx *gin.Context) {

}
func (h *AppHandler) CreateApp(ctx *gin.Context) {
	var req v1.CreateAppRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	err := h.appService.CreateApp(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	// token, err := h.adminService.Login(ctx, &req)
	// if err != nil {
	// 	v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
	// 	return
	// }
	v1.HandleSuccess(ctx, nil)
}
