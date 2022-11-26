package handler

import (
	"final-project-backend/pkg/helper"
	"final-project-backend/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PromoUserHandler struct {
	usecase usecase.PromoUserUsecase
}

func NewPromoUserHandler(usecase usecase.PromoUserUsecase) PromoUserHandler {
	return PromoUserHandler{
		usecase: usecase,
	}
}

func (h *PromoUserHandler) GetAllPromoUserByUserId(c *gin.Context) {
	userId := c.GetInt("user_id")

	resAllPromoUser, err := h.usecase.GetAllPromoUserByUserId(userId)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resAllPromoUser, http.StatusOK)
}
