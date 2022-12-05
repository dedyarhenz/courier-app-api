package handler

import (
	"final-project-backend/mocks"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPromoUserHandler_GetAllPromoUserByUserId_Error(t *testing.T) {

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/users/promos", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	s := mocks.NewPromoUserUsecase(t)
	h := NewPromoUserHandler(s)

	userId := c.GetInt("user_id")

	s.On("GetAllPromoUserByUserId", userId).Return(nil, fmt.Errorf("error"))

	h.GetAllPromoUserByUserId(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
