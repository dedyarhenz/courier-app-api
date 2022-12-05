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

func TestAddOnHandler_GetAllAddOn_Error(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/add-ons", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewAddOnUsecase(t)
	h := NewAddOnHandler(u)

	u.On("GetAllAddOn").Return(nil, fmt.Errorf("error"))

	h.GetAllAddOn(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddOnHandler_GetAllAddOn_Success(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/add-ons", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewAddOnUsecase(t)
	h := NewAddOnHandler(u)

	addOns := []dto.AddOnResponse{
		{
			Id:          1,
			Name:        "addon 1",
			Description: "desc",
			Price:       10000,
		},
		{
			Id:          2,
			Name:        "addon 2",
			Description: "desc",
			Price:       20000,
		},
	}

	u.On("GetAllAddOn").Return(addOns, nil)

	h.GetAllAddOn(c)

	assert.Equal(t, http.StatusOK, rr.Code)
}
