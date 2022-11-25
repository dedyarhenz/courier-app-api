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

type PromoHandler struct {
	usecase usecase.PromoUsecase
}

func NewPromoHandler(usecase usecase.PromoUsecase) PromoHandler {
	return PromoHandler{
		usecase: usecase,
	}
}

func (h *PromoHandler) GetAllPromo(c *gin.Context) {
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

	resAllPromo, err := h.usecase.GetAllPromo(pag, lim, search, order, sortBy)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resAllPromo, http.StatusOK)
}

func (h *PromoHandler) CreatePromo(c *gin.Context) {
	var reqPromo dto.PromoCreateRequest

	err := c.ShouldBind(&reqPromo)
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

	resPromo, err := h.usecase.CreatePromo(reqPromo)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resPromo, http.StatusCreated)
}

func (h *PromoHandler) UpdatePromo(c *gin.Context) {
	var reqPromo dto.PromoUpdateRequest

	promoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}

	err = c.ShouldBind(&reqPromo)
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

	reqPromo.Id = promoId

	resPromo, err := h.usecase.UpdatePromo(reqPromo)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resPromo, http.StatusCreated)
}
