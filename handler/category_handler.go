package handler

import (
	"final-project-backend/pkg/helper"
	"final-project-backend/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	usecase usecase.CategoryUsecase
}

func NewCategoryHandler(usecase usecase.CategoryUsecase) CategoryHandler {
	return CategoryHandler{
		usecase: usecase,
	}
}

func (h *CategoryHandler) GetAllCategory(c *gin.Context) {
	resAllCategory, err := h.usecase.GetAllCategory()
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
	}

	helper.SuccessResponse(c.Writer, resAllCategory, http.StatusOK)
}
