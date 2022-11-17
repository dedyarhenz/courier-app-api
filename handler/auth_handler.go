package handler

import (
	"final-project-backend/dto"
	"final-project-backend/pkg/helper"
	"final-project-backend/usecase"
	"fmt"
	"net/http"

	custErr "final-project-backend/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(usecase usecase.AuthUsecase) AuthHandler {
	return AuthHandler{
		usecase: usecase,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	input := dto.UserLoginRequest{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
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
		}

		helper.ErrorResponse(c.Writer, custErr.ErrInvalidRequest.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.usecase.Login(input)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, token, http.StatusCreated)
}
