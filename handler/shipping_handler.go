package handler

import (
	"final-project-backend/dto"
	custErr "final-project-backend/pkg/errors"
	"final-project-backend/pkg/helper"
	"final-project-backend/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ShippingHandler struct {
	usecase usecase.ShippingUsecase
}

func NewShippingHandler(usecase usecase.ShippingUsecase) ShippingHandler {
	return ShippingHandler{
		usecase: usecase,
	}
}

func (h *ShippingHandler) CreateShipping(c *gin.Context) {
	var reqShipping dto.ShippingCreateRequest

	err := c.ShouldBind(&reqShipping)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
			return
		}

		for _, errs := range errs {
			errMsg := fmt.Sprintf("Error field %s condition %s", errs.StructField(), errs.Tag())
			helper.ErrorResponse(c.Writer, errMsg, http.StatusBadRequest)
			return
		}
	}

	reqShipping.UserId = c.GetInt("user_id")

	resShipping, err := h.usecase.CreateShipping(reqShipping)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resShipping, http.StatusCreated)
}
