package http

import (
	"context"
	"edts-tech-test/internal/domain/entity"
	"edts-tech-test/internal/usecase"
	"edts-tech-test/internal/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type UserHandler struct {
	userUC usecase.UserUC
}

func NewUserHandler(
	userUC usecase.UserUC,
) *UserHandler {
	return &UserHandler{
		userUC: userUC,
	}
}

func (h *UserHandler) SetupHandlers(r *gin.Engine) {
	userPathV1 := r.Group("/v1")
	userPathV1.GET("/user/purchase/history/:user_phone", h.getUserPurchaseHistory)
}

func (h *UserHandler) getUserPurchaseHistory(c *gin.Context) {
	var req entity.GetUserPurchaseHistoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, utils.ErrBadRequest("Invalid Body : "+err.Error(), "UserHandler.getUserPurchaseHistory.ShouldBindJSON"))
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	resp, err := h.userUC.GetUserPurchaseHistory(ctx, &req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}
