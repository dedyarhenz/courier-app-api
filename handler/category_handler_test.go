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

func TestCategoryHandler_GetAllCategory_Error(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/categories", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewCategoryUsecase(t)
	h := NewCategoryHandler(u)

	u.On("GetAllCategory").Return(nil, fmt.Errorf("error"))

	h.GetAllCategory(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCategoryHandler_GetAllCategory_Success(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/categories", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewCategoryUsecase(t)
	h := NewCategoryHandler(u)

	categories := []dto.CategoryResponse{
		{
			Id:          1,
			Name:        "category 1",
			Description: "desc",
			Price:       10000,
		},
		{
			Id:          2,
			Name:        "category 2",
			Description: "desc",
			Price:       20000,
		},
	}

	u.On("GetAllCategory").Return(categories, nil)

	h.GetAllCategory(c)

	assert.Equal(t, http.StatusOK, rr.Code)
}
