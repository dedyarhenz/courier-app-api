package handler

import (
	"final-project-backend/pkg/helper"
	"final-project-backend/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SizeHandler struct {
	usecase usecase.SizeUsecase
}

func NewSizeHandler(usecse usecase.SizeUsecase) SizeHandler {
	return SizeHandler{
		usecase: usecse,
	}
}

func (h *SizeHandler) GetAllSize(c *gin.Context) {
	resAllSize, err := h.usecase.GetAllSize()
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resAllSize, http.StatusOK)
}
