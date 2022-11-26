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

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) UserHandler {
	return UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	userId := c.GetInt("user_id")

	resUser, err := h.usecase.GetUserById(userId)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resUser, http.StatusOK)
}

func (h *UserHandler) UpdateUserById(c *gin.Context) {
	reqUser := dto.UserUpdateRequest{}

	err := c.ShouldBind(&reqUser)
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

	file, _, err := c.Request.FormFile("photo")
	if err != nil {
		helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
	}

	reqUser.Photo = file
	reqUser.Id = c.GetInt("user_id")

	resAddress, err := h.usecase.UpdateUserById(reqUser)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resAddress, http.StatusOK)
}

func (h *UserHandler) TopUp(c *gin.Context) {
	input := dto.TopUpRequest{}
	err := c.ShouldBindJSON(&input)
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
	userId := c.GetInt("user_id")
	input.UserId = userId

	resUser, err := h.usecase.TopUp(input)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resUser, http.StatusOK)
}
