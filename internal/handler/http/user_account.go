package http

import (
	"context"
	"edts-tech-test/internal/domain/entity"
	"edts-tech-test/internal/handler/http/middleware"
	"edts-tech-test/internal/usecase"
	"edts-tech-test/internal/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type UserAccountHandler struct {
	userAccUC usecase.UserAccUC
}

func NewUserAccHandler(
	userAccUC usecase.UserAccUC,
) *UserAccountHandler {
	return &UserAccountHandler{
		userAccUC: userAccUC,
	}
}

func (h *UserAccountHandler) SetupHandlers(r *gin.Engine) {
	userAccPathV1 := r.Group("/v1")
	userAccPathV1.Use(middleware.TokenChecker)
	userAccPathV1.POST("/user-account/update-profile", h.updateProfile)
}

func (h *UserAccountHandler) updateProfile(c *gin.Context) {
	var req entity.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, utils.ErrBadRequest("Invalid Body : "+err.Error(), "TransactionHandler.updateProfile.ShouldBindJSON"))
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	ctxVal := middleware.GetContextValue(c)
	req.AccountID = ctxVal.AccountInfo.ID
	resp, err := h.userAccUC.UpdateProfile(ctx, &req)
	if err != nil {
		utils.ResponseError(c, err)
		return
	}

	utils.ResponseSuccess(c, "", resp)
}
