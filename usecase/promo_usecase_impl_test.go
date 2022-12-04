package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPromoUsecaseImpl_GetAllPromo_Error(t *testing.T) {
	mockPromoRepo := mocks.NewPromoRepository(t)
	usecase := NewPromoUsecaseImpl(mockPromoRepo)

	mockPromoRepo.On("CountPromo", "").Return(int64(1))
	mockPromoRepo.On("GetAllPromo", 0, 10, "", "expire_date DESC").Return(nil, fmt.Errorf("error"))
	res, err := usecase.GetAllPromo(1, 10, "", "expired", "DESC")

	assert.NotNil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPromoUsecaseImpl_GetAllPromo_ErrorOrder(t *testing.T) {
	mockPromoRepo := mocks.NewPromoRepository(t)
	usecase := NewPromoUsecaseImpl(mockPromoRepo)

	mockPromoRepo.On("CountPromo", "").Return(int64(1))
	mockPromoRepo.On("GetAllPromo", 0, 10, "", "expire_date DESC").Return(nil, fmt.Errorf("error"))
	res, err := usecase.GetAllPromo(1, 10, "", "expire", "DESC")

	assert.NotNil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPromoUsecaseImpl_GetAllPromo_Success(t *testing.T) {
	mockPromoRepo := mocks.NewPromoRepository(t)
	usecase := NewPromoUsecaseImpl(mockPromoRepo)

	promos := []entity.Promo{
		{
			Id:          1,
			Name:        "promo 1",
			MinFee:      10000,
			Discount:    50,
			MaxDiscount: 20000,
			Quota:       5,
		},
		{
			Id:          2,
			Name:        "promo 2",
			MinFee:      10000,
			Discount:    50,
			MaxDiscount: 20000,
			Quota:       4,
		},
	}

	promosResponse := dto.CreatePromoListResponse(promos)
	promoPaginate := dto.PromoPaginateResponse{
		Page:      1,
		Limit:     10,
		Totaldata: 2,
		TotalPage: 1,
		Data:      promosResponse,
	}

	mockPromoRepo.On("CountPromo", "").Return(int64(2))
	mockPromoRepo.On("GetAllPromo", 0, 10, "", "quota DESC").Return(promos, nil)
	res, err := usecase.GetAllPromo(1, 10, "", "quota", "DESC")

	assert.Nil(t, err)
	assert.Equal(t, promoPaginate, res)
}

func TestPromoUsecaseImpl_GetPromoById_Error(t *testing.T) {
	mockPromoRepo := mocks.NewPromoRepository(t)
	usecase := NewPromoUsecaseImpl(mockPromoRepo)

	mockPromoRepo.On("GetPromoById", 1).Return(nil, fmt.Errorf("error"))
	res, err := usecase.GetPromoById(1)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPromoUsecaseImpl_GetPromoById_Success(t *testing.T) {
	mockPromoRepo := mocks.NewPromoRepository(t)
	usecase := NewPromoUsecaseImpl(mockPromoRepo)

	promo := entity.Promo{
		Id:          1,
		Name:        "promo 1",
		MinFee:      200000,
		Discount:    50,
		MaxDiscount: 15000,
		Quota:       3,
	}

	promoResponse := dto.CreatePromoResponse(promo)

	mockPromoRepo.On("GetPromoById", 1).Return(&promo, nil)
	res, err := usecase.GetPromoById(1)

	assert.Nil(t, err)
	assert.Equal(t, &promoResponse, res)
}

func TestPromoUsecaseImpl_CreatePromo_ErrorDate(t *testing.T) {
	mockPromoRepo := mocks.NewPromoRepository(t)
	usecase := NewPromoUsecaseImpl(mockPromoRepo)

	promoRequest := dto.PromoCreateRequest{
		Name:        "promo 1",
		MinFee:      200000,
		Discount:    50,
		MaxDiscount: 15000,
		Quota:       3,
		ExpireDate:  "2022-09-1215:04:05",
	}

	res, err := usecase.CreatePromo(promoRequest)

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestPromoUsecaseImpl_CreatePromo_ErrorCreate(t *testing.T) {
	mockPromoRepo := mocks.NewPromoRepository(t)
	usecase := NewPromoUsecaseImpl(mockPromoRepo)

	promoRequest := dto.PromoCreateRequest{
		Name:        "promo 1",
		MinFee:      200000,
		Discount:    50,
		MaxDiscount: 15000,
		Quota:       3,
		ExpireDate:  "2022-09-12T15:04:05",
	}

	expiredDate, _ := time.Parse("2006-01-02T15:04:05", promoRequest.ExpireDate)

	promo := entity.Promo{
		Name:        "promo 1",
		MinFee:      200000,
		Discount:    50,
		MaxDiscount: 15000,
		Quota:       3,
		ExpireDate:  expiredDate,
	}

	mockPromoRepo.On("CreatePromo", promo).Return(nil, fmt.Errorf("error"))

	res, err := usecase.CreatePromo(promoRequest)

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestPromoUsecaseImpl_CreatePromo_Success(t *testing.T) {
	mockPromoRepo := mocks.NewPromoRepository(t)
	usecase := NewPromoUsecaseImpl(mockPromoRepo)

	promoRequest := dto.PromoCreateRequest{
		Name:        "promo 1",
		MinFee:      200000,
		Discount:    50,
		MaxDiscount: 15000,
		Quota:       3,
		ExpireDate:  "2022-09-12T15:04:05",
	}

	expiredDate, _ := time.Parse("2006-01-02T15:04:05", promoRequest.ExpireDate)

	promo := entity.Promo{
		Name:        "promo 1",
		MinFee:      200000,
		Discount:    50,
		MaxDiscount: 15000,
		Quota:       3,
		ExpireDate:  expiredDate,
	}

	promoResponse := dto.CreatePromoResponse(promo)

	mockPromoRepo.On("CreatePromo", promo).Return(&promo, nil)

	res, err := usecase.CreatePromo(promoRequest)

	assert.Nil(t, err)
	assert.Equal(t, &promoResponse, res)
}

func TestPromoUsecaseImpl_UpdatePromo_ErrorDate(t *testing.T) {
	mockPromoRepo := mocks.NewPromoRepository(t)
	usecase := NewPromoUsecaseImpl(mockPromoRepo)

	promoRequest := dto.PromoUpdateRequest{
		Id:          1,
		Name:        "promo 1",
		MinFee:      200000,
		Discount:    50,
		MaxDiscount: 15000,
		Quota:       3,
		ExpireDate:  "2022-09-1215:04:05",
	}

	res, err := usecase.UpdatePromo(promoRequest)

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestPromoUsecaseImpl_UpdatePromo_ErrorUpdate(t *testing.T) {
	mockPromoRepo := mocks.NewPromoRepository(t)
	usecase := NewPromoUsecaseImpl(mockPromoRepo)

	promoRequest := dto.PromoUpdateRequest{
		Id:          1,
		Name:        "promo 1",
		MinFee:      200000,
		Discount:    50,
		MaxDiscount: 15000,
		Quota:       3,
		ExpireDate:  "2022-09-12T15:04:05",
	}

	expiredDate, _ := time.Parse("2006-01-02T15:04:05", promoRequest.ExpireDate)

	promo := entity.Promo{
		Id:          1,
		Name:        "promo 1",
		MinFee:      200000,
		Discount:    50,
		MaxDiscount: 15000,
		Quota:       3,
		ExpireDate:  expiredDate,
	}

	mockPromoRepo.On("UpdatePromo", promo).Return(nil, fmt.Errorf("error"))

	res, err := usecase.UpdatePromo(promoRequest)

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestPromoUsecaseImpl_UpdatePromo_Success(t *testing.T) {
	mockPromoRepo := mocks.NewPromoRepository(t)
	usecase := NewPromoUsecaseImpl(mockPromoRepo)

	promoRequest := dto.PromoUpdateRequest{
		Id:          1,
		Name:        "promo 1",
		MinFee:      200000,
		Discount:    50,
		MaxDiscount: 15000,
		Quota:       3,
		ExpireDate:  "2022-09-12T15:04:05",
	}

	expiredDate, _ := time.Parse("2006-01-02T15:04:05", promoRequest.ExpireDate)

	promo := entity.Promo{
		Id:          1,
		Name:        "promo 1",
		MinFee:      200000,
		Discount:    50,
		MaxDiscount: 15000,
		Quota:       3,
		ExpireDate:  expiredDate,
	}

	promoResponse := dto.CreatePromoResponse(promo)

	mockPromoRepo.On("UpdatePromo", promo).Return(&promo, nil)

	res, err := usecase.UpdatePromo(promoRequest)

	assert.Nil(t, err)
	assert.Equal(t, &promoResponse, res)
}
