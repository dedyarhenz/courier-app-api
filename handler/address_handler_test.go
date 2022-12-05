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

func TestAddressHandler_GetAllAddress_ErrorLimit(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/addresses?limit=masndmansda", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	h.GetAllAddress(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_GetAllAddress_ErrorPage(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/addresses?limit=10&page=asdas", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	h.GetAllAddress(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_GetAllAddress_ErrorGetAllAddress(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/addresses", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Params = []gin.Param{{Key: "page", Value: "1"}, {Key: "limit", Value: "10"}, {Key: "search", Value: ""}}

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	addressPaginate := dto.AddressPaginateResponse{
		Page:      1,
		Limit:     10,
		TotalPage: 1,
		Totaldata: 2,
		Data:      []dto.AddressResponse{},
	}

	u.On("GetAllAddress", 1, 10, "").Return(addressPaginate, fmt.Errorf("error"))

	h.GetAllAddress(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_GetAllAddress_Success(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/addresses", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Params = []gin.Param{{Key: "page", Value: "1"}, {Key: "limit", Value: "10"}, {Key: "search", Value: ""}}

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	addressPaginate := dto.AddressPaginateResponse{
		Page:      1,
		Limit:     10,
		TotalPage: 1,
		Totaldata: 2,
		Data:      []dto.AddressResponse{},
	}

	u.On("GetAllAddress", 1, 10, "").Return(addressPaginate, nil)

	h.GetAllAddress(c)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestAddressHandler_GetAllAddressByUserId_ErrorLimit(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/users/addresses?limit=masndmansda", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	h.GetAllAddressByUserId(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_GetAllAddressByUserId_ErrorPage(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/users/addresses?limit=10&page=asdas", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	h.GetAllAddressByUserId(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_GetAllAddressByUserId_ErrorGetAllAddress(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/users/addresses", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Params = []gin.Param{{Key: "page", Value: "1"}, {Key: "limit", Value: "10"}, {Key: "search", Value: ""}}

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	addressPaginate := dto.AddressPaginateResponse{
		Page:      1,
		Limit:     10,
		TotalPage: 1,
		Totaldata: 2,
		Data:      []dto.AddressResponse{},
	}

	userId := c.GetInt("user_id")

	u.On("GetAllAddressByUserId", userId, 1, 10, "").Return(addressPaginate, fmt.Errorf("error"))

	h.GetAllAddressByUserId(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_GetAllAddressByUserId_Success(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/users/addresses", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Params = []gin.Param{{Key: "page", Value: "1"}, {Key: "limit", Value: "10"}, {Key: "search", Value: ""}}

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	addressPaginate := dto.AddressPaginateResponse{
		Page:      1,
		Limit:     10,
		TotalPage: 1,
		Totaldata: 2,
		Data:      []dto.AddressResponse{},
	}

	userId := c.GetInt("user_id")

	u.On("GetAllAddressByUserId", userId, 1, 10, "").Return(addressPaginate, nil)

	h.GetAllAddressByUserId(c)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestAddressHandler_GetAddressByUserId_Error(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/users/addresses/qqq", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	h.GetAddressByUserId(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_GetAddressByUserId_ErrorGetAddress(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/users/addresses/1", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: "1",
		},
	}

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	u.On("GetAddressByUserId", 0, 1).Return(nil, fmt.Errorf("error"))

	h.GetAddressByUserId(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_GetAddressByUserId_Success(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/v1/users/addresses/1", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: "1",
		},
	}

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	address := dto.AddressResponse{
		Id:             1,
		RecipientName:  "qweqr",
		FullAddress:    "asdasd",
		RecipientPhone: "asdas",
	}

	u.On("GetAddressByUserId", 0, 1).Return(&address, nil)

	h.GetAddressByUserId(c)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestAddressHandler_CreateAddress_ErrorBindJson(t *testing.T) {
	bodyReader := strings.NewReader(`{"recipient_name": "asdas91919", "full_address": "tes", "recipient_phone", "81231872"}`)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/users/addresses", bodyReader)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	h.CreateAddress(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_CreateAddress_ErrorValidation(t *testing.T) {
	bodyReader := strings.NewReader(`{}`)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/users/addresses", bodyReader)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	h.CreateAddress(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_CreateAddress_ErrorCreate(t *testing.T) {
	bodyReader := strings.NewReader(`{"recipient_name": "budi", "full_address": "tes", "recipient_phone": "81231872"}`)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/users/addresses", bodyReader)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	newAddress := dto.AddressCreateRequest{
		RecipientName:  "budi",
		FullAddress:    "tes",
		RecipientPhone: "81231872",
	}

	u.On("CreateAddress", newAddress).Return(nil, fmt.Errorf("error"))
	h.CreateAddress(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_CreateAddress_Success(t *testing.T) {
	bodyReader := strings.NewReader(`{"recipient_name": "budi", "full_address": "tes", "recipient_phone": "81231872"}`)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/v1/users/addresses", bodyReader)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	newAddress := dto.AddressCreateRequest{
		RecipientName:  "budi",
		FullAddress:    "tes",
		RecipientPhone: "81231872",
	}

	newAddressResponse := dto.AddressResponse{
		Id:             1,
		RecipientName:  "budi",
		FullAddress:    "tes",
		RecipientPhone: "81231872",
	}

	u.On("CreateAddress", newAddress).Return(&newAddressResponse, nil)
	h.CreateAddress(c)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestAddressHandler_UpdateAddress_ErrorBindJson(t *testing.T) {
	bodyReader := strings.NewReader(`{"id": "1", "recipient_name": "asdas91919", "full_address": "tes", "recipient_phone", ,"81231872"}`)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/v1/users/addresses", bodyReader)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	h.UpdateAddressByUserId(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_UpdateAddress_ErrorValidation(t *testing.T) {
	bodyReader := strings.NewReader(`{"recipient_name": "asdas91919", "full_address": "", "recipient_phone": ""}`)

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/v1/users/addresses", bodyReader)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	h.UpdateAddressByUserId(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_DeleteAddressByUserId_ErrorValidation(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/v1/users/addresses/1", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	h.DeleteAddressByUserId(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_DeleteAddressByUserId_ErrorDelete(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/v1/users/addresses", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	userId := c.GetInt("user_id")

	u.On("DeleteAddressByUserId", userId, 1).Return(fmt.Errorf("error"))

	h.DeleteAddressByUserId(c)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestAddressHandler_DeleteAddressByUserId_Success(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/v1/users/addresses", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	u := mocks.NewAddressUsecase(t)
	h := NewAddressHandler(u)

	userId := c.GetInt("user_id")

	u.On("DeleteAddressByUserId", userId, 1).Return(nil)

	h.DeleteAddressByUserId(c)

	assert.Equal(t, http.StatusOK, rr.Code)
}
