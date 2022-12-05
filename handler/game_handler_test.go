package handler

import (
	"final-project-backend/dto"
	"final-project-backend/mocks"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGameHandler_Play_ErrorBinJson(t *testing.T) {
	bodyReader := strings.NewReader(`{"shipping_id": "adads"}`)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/users/games/play", bodyReader)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewGameUsecase(t)
	h := NewGameHandler(u)

	h.Play(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGameHandler_Play_ErrorValidation(t *testing.T) {
	bodyReader := strings.NewReader(`{}`)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/users/games/play", bodyReader)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewGameUsecase(t)
	h := NewGameHandler(u)

	h.Play(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGameHandler_Play_ErrorPlay(t *testing.T) {
	bodyReader := strings.NewReader(`{"shipping_id": 1}`)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/users/games/play", bodyReader)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewGameUsecase(t)
	h := NewGameHandler(u)

	u.On("Play", dto.GamePlayRequest{ShippingId: 1, UserId: 0}).Return(nil, fmt.Errorf("error"))

	h.Play(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGameHandler_Play_Success(t *testing.T) {
	bodyReader := strings.NewReader(`{"shipping_id": 1}`)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/users/games/play", bodyReader)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewGameUsecase(t)
	h := NewGameHandler(u)

	gameResponse := dto.GameResponse{
		Name:        "tes promo",
		MinFee:      2000,
		Discount:    20,
		MaxDiscount: 20000,
	}

	u.On("Play", dto.GamePlayRequest{ShippingId: 1, UserId: 0}).Return(&gameResponse, nil)

	h.Play(c)

	assert.Equal(t, http.StatusOK, rr.Code)
}
