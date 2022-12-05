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

type GameHandler struct {
	usecase usecase.GameUsecase
}

func NewGameHandler(usecase usecase.GameUsecase) GameHandler {
	return GameHandler{
		usecase: usecase,
	}
}

func (h *GameHandler) Play(c *gin.Context) {
	reqGamePlay := dto.GamePlayRequest{}

	err := c.ShouldBindJSON(&reqGamePlay)
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

	reqGamePlay.UserId = c.GetInt("user_id")

	resGame, err := h.usecase.Play(reqGamePlay)
	if err != nil {
		helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	helper.SuccessResponse(c.Writer, resGame, http.StatusOK)
}
