package handler

import (
	"final-project-backend/pkg/helper"
	"final-project-backend/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddOnHandler struct {
	usecase usecase.AddOnUsecase
}

func NewAddOnHandler(usecase usecase.AddOnUsecase) AddOnHandler {
	return AddOnHandler{
		usecase: usecase,
	}
}

func (h *AddOnHandler) GetAllAddOn(c *gin.Context) {
	resAllAddOn, err := h.usecase.GetAllAddOn()
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
	}

	helper.SuccessResponse(c.Writer, resAllAddOn, http.StatusOK)
}
