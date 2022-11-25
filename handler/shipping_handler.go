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

func (h *ShippingHandler) GetAllShipping(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "1")
	order := c.DefaultQuery("order", "expired")
	sortBy := c.DefaultQuery("sortBy", "desc")

	lim, err := strconv.Atoi(limit)
	if err != nil {
		helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}

	pag, err := strconv.Atoi(page)
	if err != nil {
		helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}

	resShippings, err := h.usecase.GetAllShipping(pag, lim, search, order, sortBy)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resShippings, http.StatusOK)
}

func (h *ShippingHandler) GetAllShippingByUserId(c *gin.Context) {
	userId := c.GetInt("user_id")
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "1")
	order := c.DefaultQuery("order", "date")
	sortBy := c.DefaultQuery("sortBy", "desc")

	lim, err := strconv.Atoi(limit)
	if err != nil {
		helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}

	pag, err := strconv.Atoi(page)
	if err != nil {
		helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}

	resShippings, err := h.usecase.GetAllShippingByUserId(userId, pag, lim, search, order, sortBy)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resShippings, http.StatusOK)
}

func (h *ShippingHandler) GetShippingByUserId(c *gin.Context) {
	shippingId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}
	userId := c.GetInt("user_id")

	resShippings, err := h.usecase.GetShippingByUserId(userId, shippingId)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resShippings, http.StatusOK)
}

func (h *ShippingHandler) UpdateReviewByUserId(c *gin.Context) {

	shippingId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}

	reqShippingReview := dto.ShippingReviewRequest{
		UserId:     c.GetInt("user_id"),
		ShippingId: shippingId,
	}

	err = c.ShouldBind(&reqShippingReview)
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

	err = h.usecase.UpdateReviewByUserId(reqShippingReview)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, nil, http.StatusOK)
}
