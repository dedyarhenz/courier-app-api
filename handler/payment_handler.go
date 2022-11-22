package handler

import (
	"final-project-backend/dto"
	custErr "final-project-backend/pkg/errors"
	"final-project-backend/pkg/helper"
	"final-project-backend/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentHandler struct {
	usecase usecase.PaymenUsecase
}

func NewPaymentHandler(usecase usecase.PaymenUsecase) PaymentHandler {
	return PaymentHandler{
		usecase: usecase,
	}
}

func (h *PaymentHandler) PayUserShipping(c *gin.Context) {
	var reqPaymentPay dto.PaymentPayRequest
	err := c.ShouldBindJSON(&reqPaymentPay)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
			return
		}

		for _, errs := range errs {
			errMsg := fmt.Sprintf("Error field %s condition %s", errs.Field(), errs.Tag())
			helper.ErrorResponse(c.Writer, errMsg, http.StatusBadRequest)
			return
		}

		helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}

	paymentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}

	reqPaymentPay.PaymentId = paymentId
	reqPaymentPay.UserId = c.GetInt("user_id")

	resPayment, err := h.usecase.PayUserShipping(reqPaymentPay)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resPayment, http.StatusOK)
}
