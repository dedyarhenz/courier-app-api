package handler

import (
	"final-project-backend/dto"
	"final-project-backend/mocks"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSizeHandler_GetAllSize_Error(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/sizes", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewSizeUsecase(t)
	h := NewSizeHandler(u)

	u.On("GetAllSize").Return(nil, fmt.Errorf("error"))

	h.GetAllSize(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestSizeHandler_GetAllSize_Success(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/sizes", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewSizeUsecase(t)
	h := NewSizeHandler(u)

	sizes := []dto.SizeResponse{
		{
			Id:          1,
			Name:        "size 1",
			Description: "desc",
			Price:       10000,
		},
		{
			Id:          2,
			Name:        "size 2",
			Description: "desc",
			Price:       20000,
		},
	}

	u.On("GetAllSize").Return(sizes, nil)

	h.GetAllSize(c)

	assert.Equal(t, http.StatusOK, rr.Code)
}
