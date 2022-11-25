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

type AddressHandler struct {
	usecase usecase.AddressUsecase
}

func NewAddressHandler(usecase usecase.AddressUsecase) AddressHandler {
	return AddressHandler{
		usecase: usecase,
	}
}

func (h *AddressHandler) GetAllAddress(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "1")

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

	resAllAddress, err := h.usecase.GetAllAddress(pag, lim, search)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resAllAddress, http.StatusOK)
}

func (h *AddressHandler) GetAllAddressByUserId(c *gin.Context) {
	userId := c.GetInt("user_id")
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "1")

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

	resAllAddress, err := h.usecase.GetAllAddressByUserId(userId, pag, lim, search)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resAllAddress, http.StatusOK)
}

func (h *AddressHandler) GetAddressByUserId(c *gin.Context) {
	addressId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}
	userId := c.GetInt("user_id")

	resAddress, err := h.usecase.GetAddressByUserId(userId, addressId)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resAddress, http.StatusOK)
}

func (h *AddressHandler) CreateAddress(c *gin.Context) {
	var reqAddress dto.AddressCreateRequest

	err := c.ShouldBind(&reqAddress)
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

	reqAddress.UserId = c.GetInt("user_id")

	resAddress, err := h.usecase.CreateAddress(reqAddress)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resAddress, http.StatusCreated)
}

func (h *AddressHandler) UpdateAddressByUserId(c *gin.Context) {
	var reqAddress dto.AddressUpdateRequest

	err := c.ShouldBind(&reqAddress)
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

	addressId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}

	reqAddress.UserId = c.GetInt("user_id")
	reqAddress.Id = addressId

	resAddress, err := h.usecase.UpdateAddressByUserId(reqAddress)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resAddress, http.StatusOK)
}
