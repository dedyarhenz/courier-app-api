package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	custErr "final-project-backend/pkg/errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShippingUsecaseImpl_GetAllShipping_Error(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	mockShippingRepo.On("CountShipping", "").Return(int64(11))
	mockShippingRepo.On("GetAllShipping", 0, 10, "", "created_at DESC").Return(nil, fmt.Errorf("error"))
	res, err := usecase.GetAllShipping(1, 10, "", "date", "DESC")

	assert.NotNil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_GetAllShipping_Success(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	shippings := []entity.Shipping{
		{
			Id:             1,
			SizeId:         1,
			CategoryId:     1,
			AddressId:      1,
			PaymentId:      1,
			StatusShipping: "PROCESS",
			Review:         nil,
			IsPlayGame:     false,
		},
		{
			Id:             2,
			SizeId:         2,
			CategoryId:     2,
			AddressId:      2,
			PaymentId:      2,
			StatusShipping: "PROCESS",
			Review:         nil,
			IsPlayGame:     false,
		},
	}

	shippingsResponse := dto.CreateShippingListResponse(shippings)
	shippingPaginate := dto.ShippingPaginateResponse{
		Page:      1,
		Limit:     10,
		Totaldata: 2,
		TotalPage: 0,
		Data:      shippingsResponse,
	}

	tests := []struct {
		order    string
		orderRes string
	}{
		{
			order:    "date",
			orderRes: "created_at",
		},
		{
			order:    "category",
			orderRes: "category_id",
		},
		{
			order:    "size",
			orderRes: "size_id",
		},
		{
			order:    "payment",
			orderRes: "payments.total_cost",
		},
		{
			order:    "status",
			orderRes: "status_shipping",
		},
		{
			order:    "asdad",
			orderRes: "created_at",
		},
	}

	for _, tt := range tests {
		t.Run(tt.order, func(t *testing.T) {

			mockShippingRepo.On("CountShipping", "").Return(int64(2))
			mockShippingRepo.On("GetAllShipping", 0, 10, "", fmt.Sprintf("%s DESC", tt.orderRes)).Return(shippings, nil)

			res, err := usecase.GetAllShipping(1, 10, "", tt.order, "DESC")

			assert.Nil(t, err)
			assert.Equal(t, shippingPaginate, res)
		})
	}
}

func TestShippingUsecaseImpl_GetAllReportShippingByDate_Error(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	startDate := time.Date(2022, time.Month(1), 01, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
	endDate := time.Date(2022, time.Month(1), 31, 0, 0, 0, 0, time.UTC).Format("2006-01-02")

	mockShippingRepo.On("CountShippingByDate", startDate, endDate).Return(int64(11))
	mockShippingRepo.On("GetAllReportShippingByDate", startDate, endDate, 0, 10, "created_at DESC").Return(nil, fmt.Errorf("error"))
	res, err := usecase.GetAllReportShippingByDate(1, 2022, 1, 10, "DESC")

	assert.NotNil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_GetAllReportShippingByDate_Success(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	startDate := time.Date(2022, time.Month(1), 01, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
	endDate := time.Date(2022, time.Month(1), 31, 0, 0, 0, 0, time.UTC).Format("2006-01-02")

	shippings := []entity.Shipping{
		{
			Id:             1,
			SizeId:         1,
			CategoryId:     1,
			AddressId:      1,
			PaymentId:      1,
			StatusShipping: "PROCESS",
			Review:         nil,
			IsPlayGame:     false,
		},
		{
			Id:             2,
			SizeId:         2,
			CategoryId:     2,
			AddressId:      2,
			PaymentId:      2,
			StatusShipping: "PROCESS",
			Review:         nil,
			IsPlayGame:     false,
		},
	}

	shippingsResponse := dto.CreateShippingListResponse(shippings)
	shippingPaginate := dto.ShippingReportPaginateResponse{
		Page:           1,
		Limit:          10,
		Totaldata:      2,
		TotalPage:      0,
		Data:           shippingsResponse,
		TotalCostMonth: 20000000,
	}

	mockShippingRepo.On("CountShippingByDate", startDate, endDate).Return(int64(2))
	mockShippingRepo.On("GetAllReportShippingByDate", startDate, endDate, 0, 10, fmt.Sprintf("%s DESC", "created_at")).Return(shippings, nil)
	mockPaymentRepo.On("TotalCostPaymentByDate", startDate, endDate).Return(int64(20000000))
	res, err := usecase.GetAllReportShippingByDate(1, 2022, 1, 10, "DESC")

	assert.Nil(t, err)
	assert.Equal(t, shippingPaginate, res)
}

func TestShippingUsecaseImpl_GetShippingById_Error(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	mockShippingRepo.On("GetShippingById", 1).Return(nil, fmt.Errorf("error"))
	res, err := usecase.GetShippingById(1)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_GetShippingById_Success(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	shipping := entity.Shipping{
		Id:             1,
		SizeId:         1,
		CategoryId:     1,
		AddressId:      1,
		PaymentId:      1,
		StatusShipping: "PROCESS",
		Review:         nil,
		IsPlayGame:     false,
	}

	shippingResponse := dto.CreateShippingResponse(shipping)

	mockShippingRepo.On("GetShippingById", 1).Return(&shipping, nil)
	res, err := usecase.GetShippingById(1)

	assert.Nil(t, err)
	assert.Equal(t, &shippingResponse, res)
}

func TestShippingUsecaseImpl_GetAllShippingByUserId_Error(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	userId := 1

	mockShippingRepo.On("CountShippingByUserId", userId, "").Return(int64(1))
	mockShippingRepo.On("GetAllShippingByUserId", userId, 0, 10, "", "created_at DESC").Return(nil, fmt.Errorf("error"))
	res, err := usecase.GetAllShippingByUserId(userId, 1, 10, "", "date", "DESC")

	assert.NotNil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_GetAllShippingByUserId_Success(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	userId := 1

	shippings := []entity.Shipping{
		{
			Id:             1,
			SizeId:         1,
			CategoryId:     1,
			AddressId:      1,
			PaymentId:      1,
			StatusShipping: "PROCESS",
			Review:         nil,
			IsPlayGame:     false,
		},
		{
			Id:             2,
			SizeId:         2,
			CategoryId:     2,
			AddressId:      2,
			PaymentId:      2,
			StatusShipping: "PROCESS",
			Review:         nil,
			IsPlayGame:     false,
		},
	}

	shippingsResponse := dto.CreateShippingListResponse(shippings)
	shippingPaginate := dto.ShippingPaginateResponse{
		Page:      1,
		Limit:     10,
		Totaldata: 2,
		TotalPage: 1,
		Data:      shippingsResponse,
	}

	mockShippingRepo.On("CountShippingByUserId", userId, "").Return(int64(2))
	mockShippingRepo.On("GetAllShippingByUserId", userId, 0, 10, "", "created_at DESC").Return(shippings, nil)
	res, err := usecase.GetAllShippingByUserId(userId, 1, 10, "", "date", "DESC")

	assert.Nil(t, err)
	assert.Equal(t, shippingPaginate, res)
}

func TestShippingUsecaseImpl_GetShippingByUserId_Error(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	mockShippingRepo.On("GetShippingByUserId", 1, 1).Return(nil, fmt.Errorf("error"))
	res, err := usecase.GetShippingByUserId(1, 1)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_GetShippingByUserId_Success(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	shipping := entity.Shipping{
		Id:             1,
		SizeId:         1,
		CategoryId:     1,
		AddressId:      1,
		PaymentId:      1,
		StatusShipping: "PROCESS",
		Review:         nil,
		IsPlayGame:     false,
	}

	shippingResponse := dto.CreateShippingResponse(shipping)

	mockShippingRepo.On("GetShippingByUserId", 1, 1).Return(&shipping, nil)
	res, err := usecase.GetShippingByUserId(1, 1)

	assert.Nil(t, err)
	assert.Equal(t, &shippingResponse, res)
}

func TestShippingUsecaseImpl_UpdateReviewByUserId_ErrorShippingNotFound(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	reviewRequest := dto.ShippingReviewRequest{
		UserId:     1,
		ShippingId: 1,
		Review:     "Ok",
	}

	mockShippingRepo.On("GetShippingByUserId", reviewRequest.UserId, reviewRequest.ShippingId).Return(nil, fmt.Errorf("error"))

	err := usecase.UpdateReviewByUserId(reviewRequest)

	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_UpdateReviewByUserId_ErrorShippingAlreadyReview(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	reviewRequest := dto.ShippingReviewRequest{
		UserId:     1,
		ShippingId: 1,
		Review:     "Ok",
	}

	review := "sudah review"
	shipping := entity.Shipping{
		Id:             1,
		SizeId:         1,
		CategoryId:     1,
		AddressId:      1,
		PaymentId:      1,
		StatusShipping: "PROCESS",
		Review:         &review,
		IsPlayGame:     false,
	}

	mockShippingRepo.On("GetShippingByUserId", reviewRequest.UserId, reviewRequest.ShippingId).Return(&shipping, nil)

	err := usecase.UpdateReviewByUserId(reviewRequest)

	assert.Equal(t, custErr.ErrShippingAlreadyReview, err)
}

func TestShippingUsecaseImpl_UpdateReviewByUserId_ErrorShippingNotDone(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	reviewRequest := dto.ShippingReviewRequest{
		UserId:     1,
		ShippingId: 1,
		Review:     "Ok",
	}

	shipping := entity.Shipping{
		Id:             1,
		SizeId:         1,
		CategoryId:     1,
		AddressId:      1,
		PaymentId:      1,
		StatusShipping: "PROCESS",
		Review:         nil,
		IsPlayGame:     false,
		Payment: &entity.Payment{
			PaymentStatus: "PENDING",
		},
	}

	t.Run("status shipping not done", func(t *testing.T) {
		mockShippingRepo.On("GetShippingByUserId", reviewRequest.UserId, reviewRequest.ShippingId).Return(&shipping, nil)
		err := usecase.UpdateReviewByUserId(reviewRequest)
		assert.Equal(t, custErr.ErrShippingReview, err)
	})

	t.Run("status payment not done", func(t *testing.T) {
		shipping.StatusShipping = "DELIVERED"

		mockShippingRepo.On("GetShippingByUserId", reviewRequest.UserId, reviewRequest.ShippingId).Return(&shipping, nil)
		err := usecase.UpdateReviewByUserId(reviewRequest)
		assert.Equal(t, custErr.ErrShippingReview, err)
	})
}

func TestShippingUsecaseImpl_UpdateReviewByUserId_ErrorShippingUdateReview(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	reviewRequest := dto.ShippingReviewRequest{
		UserId:     1,
		ShippingId: 1,
		Review:     "Ok",
	}

	shipping := entity.Shipping{
		Id:             1,
		SizeId:         1,
		CategoryId:     1,
		AddressId:      1,
		PaymentId:      1,
		StatusShipping: "DELIVERED",
		Review:         nil,
		IsPlayGame:     false,
		Payment: &entity.Payment{
			PaymentStatus: "SUCCESS",
		},
	}

	mockShippingRepo.On("GetShippingByUserId", reviewRequest.UserId, reviewRequest.ShippingId).Return(&shipping, nil)
	mockShippingRepo.On("UpdateReviewByUserId", reviewRequest.UserId, reviewRequest.ShippingId, reviewRequest.Review).Return(fmt.Errorf("error"))

	err := usecase.UpdateReviewByUserId(reviewRequest)

	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_UpdateReviewByUserId_Success(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	reviewRequest := dto.ShippingReviewRequest{
		UserId:     1,
		ShippingId: 1,
		Review:     "Ok",
	}

	shipping := entity.Shipping{
		Id:             1,
		SizeId:         1,
		CategoryId:     1,
		AddressId:      1,
		PaymentId:      1,
		StatusShipping: "DELIVERED",
		Review:         nil,
		IsPlayGame:     false,
		Payment: &entity.Payment{
			PaymentStatus: "SUCCESS",
		},
	}

	mockShippingRepo.On("GetShippingByUserId", reviewRequest.UserId, reviewRequest.ShippingId).Return(&shipping, nil)
	mockShippingRepo.On("UpdateReviewByUserId", reviewRequest.UserId, reviewRequest.ShippingId, reviewRequest.Review).Return(nil)

	err := usecase.UpdateReviewByUserId(reviewRequest)

	assert.Equal(t, nil, err)
}

func TestShippingUsecaseImpl_UpdateStatusShipping_ErrorShippingStatus(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	updateStatusRequest := dto.ShippingUpdateStatusRequest{
		ShippingId:     1,
		StatusShipping: "asdas",
	}

	err := usecase.UpdateStatusShipping(updateStatusRequest)

	assert.Equal(t, custErr.ErrShippingStatus, err)
}

func TestShippingUsecaseImpl_UpdateStatusShipping_ErrorShippingNotFound(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	updateStatusRequest := dto.ShippingUpdateStatusRequest{
		ShippingId:     1,
		StatusShipping: "PROCESS",
	}

	mockShippingRepo.On("GetShippingById", updateStatusRequest.ShippingId).Return(nil, fmt.Errorf("error"))

	err := usecase.UpdateStatusShipping(updateStatusRequest)

	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_UpdateStatusShipping_ErrorPaymentPending(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	updateStatusRequest := dto.ShippingUpdateStatusRequest{
		ShippingId:     1,
		StatusShipping: "PROCESS",
	}

	shipping := entity.Shipping{
		Id:             1,
		SizeId:         1,
		CategoryId:     1,
		AddressId:      1,
		PaymentId:      1,
		StatusShipping: "PROCESS",
		Review:         nil,
		IsPlayGame:     false,
		Payment: &entity.Payment{
			PaymentStatus: "PENDING",
		},
	}

	mockShippingRepo.On("GetShippingById", updateStatusRequest.ShippingId).Return(&shipping, nil)

	err := usecase.UpdateStatusShipping(updateStatusRequest)

	assert.Equal(t, custErr.ErrShippingMustPaid, err)
}

func TestShippingUsecaseImpl_UpdateStatusShipping_ErrorUpdateStatus(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	updateStatusRequest := dto.ShippingUpdateStatusRequest{
		ShippingId:     1,
		StatusShipping: "PROCESS",
	}

	shipping := entity.Shipping{
		Id:             1,
		SizeId:         1,
		CategoryId:     1,
		AddressId:      1,
		PaymentId:      1,
		StatusShipping: "PROCESS",
		Review:         nil,
		IsPlayGame:     false,
		Payment: &entity.Payment{
			PaymentStatus: "SUCCESS",
		},
	}

	mockShippingRepo.On("GetShippingById", updateStatusRequest.ShippingId).Return(&shipping, nil)
	mockShippingRepo.On("UpdateStatusShipping", updateStatusRequest.ShippingId, updateStatusRequest.StatusShipping).Return(fmt.Errorf("error"))

	err := usecase.UpdateStatusShipping(updateStatusRequest)

	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_UpdateStatusShipping_Success(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	updateStatusRequest := dto.ShippingUpdateStatusRequest{
		ShippingId:     1,
		StatusShipping: "PROCESS",
	}

	shipping := entity.Shipping{
		Id:             1,
		SizeId:         1,
		CategoryId:     1,
		AddressId:      1,
		PaymentId:      1,
		StatusShipping: "PROCESS",
		Review:         nil,
		IsPlayGame:     false,
		Payment: &entity.Payment{
			PaymentStatus: "SUCCESS",
		},
	}

	tests := []struct {
		statusRequest string
		status        string
	}{
		{
			statusRequest: "process",
			status:        "PROCESS",
		},
		{
			statusRequest: "pickup",
			status:        "PICKUP",
		},
		{
			statusRequest: "delivery",
			status:        "DELIVERY",
		},
		{
			statusRequest: "delivered",
			status:        "DELIVERED",
		},
	}

	for _, tt := range tests {
		t.Run(tt.statusRequest, func(t *testing.T) {
			updateStatusRequest.StatusShipping = tt.status
			mockShippingRepo.On("GetShippingById", updateStatusRequest.ShippingId).Return(&shipping, nil)
			mockShippingRepo.On("UpdateStatusShipping", updateStatusRequest.ShippingId, updateStatusRequest.StatusShipping).Return(nil)

			err := usecase.UpdateStatusShipping(updateStatusRequest)

			assert.Equal(t, nil, err)
		})
	}
}

func TestShippingUsecaseImpl_CreateShipping_ErrorNoAddOn(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	shippingCreateRequest := dto.ShippingCreateRequest{
		UserId:     1,
		SizeId:     1,
		CategoryId: 1,
		AddressId:  1,
		AddOnsId:   []int{},
	}

	res, err := usecase.CreateShipping(shippingCreateRequest)

	assert.Nil(t, res)
	assert.Equal(t, custErr.ErrMinAddOns, err)
}

func TestShippingUsecaseImpl_CreateShipping_ErrorDuplicateAddOn(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	shippingCreateRequest := dto.ShippingCreateRequest{
		UserId:     1,
		SizeId:     1,
		CategoryId: 1,
		AddressId:  1,
		AddOnsId:   []int{1, 1},
	}

	mockAddOnRepo.On("GetAddOnByMultipleId", shippingCreateRequest.AddOnsId).Return(nil, fmt.Errorf("error"))

	res, err := usecase.CreateShipping(shippingCreateRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_CreateShipping_ErrorSizeNotFound(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	shippingCreateRequest := dto.ShippingCreateRequest{
		UserId:     1,
		SizeId:     1,
		CategoryId: 1,
		AddressId:  1,
		AddOnsId:   []int{1, 2},
	}

	addOns := []entity.AddOn{
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

	mockAddOnRepo.On("GetAddOnByMultipleId", shippingCreateRequest.AddOnsId).Return(addOns, nil)
	mockSizeRepo.On("GetSizeById", shippingCreateRequest.SizeId).Return(nil, fmt.Errorf("error"))

	res, err := usecase.CreateShipping(shippingCreateRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_CreateShipping_ErrorCategoryNotFound(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	shippingCreateRequest := dto.ShippingCreateRequest{
		UserId:     1,
		SizeId:     1,
		CategoryId: 1,
		AddressId:  1,
		AddOnsId:   []int{1, 2},
	}

	addOns := []entity.AddOn{
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

	size := entity.Size{
		Id:          1,
		Name:        "size",
		Description: "desc",
		Price:       20000,
	}

	mockAddOnRepo.On("GetAddOnByMultipleId", shippingCreateRequest.AddOnsId).Return(addOns, nil)
	mockSizeRepo.On("GetSizeById", shippingCreateRequest.SizeId).Return(&size, nil)
	mockCategoryRepo.On("GetCategoryById", shippingCreateRequest.CategoryId).Return(nil, fmt.Errorf("error"))

	res, err := usecase.CreateShipping(shippingCreateRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_CreateShipping_ErrorUserNotFound(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	shippingCreateRequest := dto.ShippingCreateRequest{
		UserId:     1,
		SizeId:     1,
		CategoryId: 1,
		AddressId:  1,
		AddOnsId:   []int{1, 2},
	}

	addOns := []entity.AddOn{
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

	size := entity.Size{
		Id:          1,
		Name:        "size",
		Description: "desc",
		Price:       20000,
	}

	category := entity.Category{
		Id:          1,
		Name:        "category",
		Description: "desc",
		Price:       20000,
	}

	mockAddOnRepo.On("GetAddOnByMultipleId", shippingCreateRequest.AddOnsId).Return(addOns, nil)
	mockSizeRepo.On("GetSizeById", shippingCreateRequest.SizeId).Return(&size, nil)
	mockCategoryRepo.On("GetCategoryById", shippingCreateRequest.CategoryId).Return(&category, nil)
	mockUserRepo.On("GetUserById", shippingCreateRequest.UserId).Return(nil, fmt.Errorf("error"))

	res, err := usecase.CreateShipping(shippingCreateRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_CreateShipping_ErrorAddressNotFound(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	shippingCreateRequest := dto.ShippingCreateRequest{
		UserId:     1,
		SizeId:     1,
		CategoryId: 1,
		AddressId:  1,
		AddOnsId:   []int{1, 2},
	}

	addOns := []entity.AddOn{
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

	size := entity.Size{
		Id:          1,
		Name:        "size",
		Description: "desc",
		Price:       20000,
	}

	category := entity.Category{
		Id:          1,
		Name:        "category",
		Description: "desc",
		Price:       20000,
	}

	user := entity.User{
		Id:    1,
		Email: "dedy@gmail.com",
	}

	mockAddOnRepo.On("GetAddOnByMultipleId", shippingCreateRequest.AddOnsId).Return(addOns, nil)
	mockSizeRepo.On("GetSizeById", shippingCreateRequest.SizeId).Return(&size, nil)
	mockCategoryRepo.On("GetCategoryById", shippingCreateRequest.CategoryId).Return(&category, nil)
	mockUserRepo.On("GetUserById", shippingCreateRequest.UserId).Return(&user, nil)
	mockAddressRepo.On("GetAddressByUserId", shippingCreateRequest.UserId, shippingCreateRequest.AddressId).Return(nil, fmt.Errorf("error"))

	res, err := usecase.CreateShipping(shippingCreateRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_CreateShipping_ErrorTotalAddOnNotMatch(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	shippingCreateRequest := dto.ShippingCreateRequest{
		UserId:     1,
		SizeId:     1,
		CategoryId: 1,
		AddressId:  1,
		AddOnsId:   []int{1},
	}

	addOns := []entity.AddOn{
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

	size := entity.Size{
		Id:          1,
		Name:        "size",
		Description: "desc",
		Price:       20000,
	}

	category := entity.Category{
		Id:          1,
		Name:        "category",
		Description: "desc",
		Price:       20000,
	}

	user := entity.User{
		Id:    1,
		Email: "dedy@gmail.com",
	}

	address := entity.Address{
		Id:             1,
		RecipientName:  "budi",
		FullAddress:    "jakarta",
		RecipientPhone: "098",
		UserId:         1,
	}

	mockAddOnRepo.On("GetAddOnByMultipleId", shippingCreateRequest.AddOnsId).Return(addOns, nil)
	mockSizeRepo.On("GetSizeById", shippingCreateRequest.SizeId).Return(&size, nil)
	mockCategoryRepo.On("GetCategoryById", shippingCreateRequest.CategoryId).Return(&category, nil)
	mockUserRepo.On("GetUserById", shippingCreateRequest.UserId).Return(&user, nil)
	mockAddressRepo.On("GetAddressByUserId", shippingCreateRequest.UserId, shippingCreateRequest.AddressId).Return(&address, nil)

	res, err := usecase.CreateShipping(shippingCreateRequest)

	assert.Nil(t, res)
	assert.Equal(t, custErr.ErrAddOnInvalid, err)
}

func TestShippingUsecaseImpl_CreateShipping_ErrorCreatePayment(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	shippingCreateRequest := dto.ShippingCreateRequest{
		UserId:     1,
		SizeId:     1,
		CategoryId: 1,
		AddressId:  1,
		AddOnsId:   []int{1, 2},
	}

	addOns := []entity.AddOn{
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

	size := entity.Size{
		Id:          1,
		Name:        "size",
		Description: "desc",
		Price:       20000,
	}

	category := entity.Category{
		Id:          1,
		Name:        "category",
		Description: "desc",
		Price:       20000,
	}

	user := entity.User{
		Id:    1,
		Email: "dedy@gmail.com",
	}

	address := entity.Address{
		Id:             1,
		RecipientName:  "budi",
		FullAddress:    "jakarta",
		RecipientPhone: "098",
		UserId:         1,
	}

	payment := entity.Payment{
		PaymentStatus: "PENDING",
		TotalCost:     70000,
	}

	mockAddOnRepo.On("GetAddOnByMultipleId", shippingCreateRequest.AddOnsId).Return(addOns, nil)
	mockSizeRepo.On("GetSizeById", shippingCreateRequest.SizeId).Return(&size, nil)
	mockCategoryRepo.On("GetCategoryById", shippingCreateRequest.CategoryId).Return(&category, nil)
	mockUserRepo.On("GetUserById", shippingCreateRequest.UserId).Return(&user, nil)
	mockAddressRepo.On("GetAddressByUserId", shippingCreateRequest.UserId, shippingCreateRequest.AddressId).Return(&address, nil)
	mockPaymentRepo.On("CreatePayment", payment).Return(nil, fmt.Errorf("error"))

	res, err := usecase.CreateShipping(shippingCreateRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_CreateShipping_ErrorCreateShipping(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	shippingCreateRequest := dto.ShippingCreateRequest{
		UserId:     1,
		SizeId:     1,
		CategoryId: 1,
		AddressId:  1,
		AddOnsId:   []int{1, 2},
	}

	addOns := []entity.AddOn{
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

	size := entity.Size{
		Id:          1,
		Name:        "size",
		Description: "desc",
		Price:       20000,
	}

	category := entity.Category{
		Id:          1,
		Name:        "category",
		Description: "desc",
		Price:       20000,
	}

	user := entity.User{
		Id:    1,
		Email: "dedy@gmail.com",
	}

	address := entity.Address{
		Id:             1,
		RecipientName:  "budi",
		FullAddress:    "jakarta",
		RecipientPhone: "098",
		UserId:         1,
	}

	payment := entity.Payment{
		PaymentStatus: "PENDING",
		TotalCost:     70000,
	}

	newAddOnShipping := []entity.AddOnShipping{
		{
			AddOnId: 1,
		},
		{
			AddOnId: 2,
		},
	}

	shipping := entity.Shipping{
		SizeId:         size.Id,
		CategoryId:     category.Id,
		AddressId:      address.Id,
		PaymentId:      payment.Id,
		AddOnShippings: newAddOnShipping,
		StatusShipping: "PROCESS",
	}

	mockAddOnRepo.On("GetAddOnByMultipleId", shippingCreateRequest.AddOnsId).Return(addOns, nil)
	mockSizeRepo.On("GetSizeById", shippingCreateRequest.SizeId).Return(&size, nil)
	mockCategoryRepo.On("GetCategoryById", shippingCreateRequest.CategoryId).Return(&category, nil)
	mockUserRepo.On("GetUserById", shippingCreateRequest.UserId).Return(&user, nil)
	mockAddressRepo.On("GetAddressByUserId", shippingCreateRequest.UserId, shippingCreateRequest.AddressId).Return(&address, nil)
	mockPaymentRepo.On("CreatePayment", payment).Return(&payment, nil)
	mockShippingRepo.On("CreateShipping", shipping).Return(nil, fmt.Errorf("error"))

	res, err := usecase.CreateShipping(shippingCreateRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestShippingUsecaseImpl_CreateShipping_Success(t *testing.T) {
	mockShippingRepo := mocks.NewShippingRepository(t)
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockAddressRepo := mocks.NewAddressRepository(t)
	mockSizeRepo := mocks.NewSizeRepository(t)
	mockCategoryRepo := mocks.NewCategoryRepository(t)
	mockAddOnRepo := mocks.NewAddOnRepository(t)
	usecase := NewShippingUsecaseImpl(mockShippingRepo, mockPaymentRepo, mockUserRepo, mockAddressRepo, mockSizeRepo, mockCategoryRepo, mockAddOnRepo)

	shippingCreateRequest := dto.ShippingCreateRequest{
		UserId:     1,
		SizeId:     1,
		CategoryId: 1,
		AddressId:  1,
		AddOnsId:   []int{1, 2},
	}

	addOns := []entity.AddOn{
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

	size := entity.Size{
		Id:          1,
		Name:        "size",
		Description: "desc",
		Price:       20000,
	}

	category := entity.Category{
		Id:          1,
		Name:        "category",
		Description: "desc",
		Price:       20000,
	}

	user := entity.User{
		Id:    1,
		Email: "dedy@gmail.com",
	}

	address := entity.Address{
		Id:             1,
		RecipientName:  "budi",
		FullAddress:    "jakarta",
		RecipientPhone: "098",
		UserId:         1,
	}

	payment := entity.Payment{
		PaymentStatus: "PENDING",
		TotalCost:     70000,
	}

	newAddOnShipping := []entity.AddOnShipping{
		{
			AddOnId: 1,
		},
		{
			AddOnId: 2,
		},
	}

	shipping := entity.Shipping{
		SizeId:         size.Id,
		CategoryId:     category.Id,
		AddressId:      address.Id,
		PaymentId:      payment.Id,
		AddOnShippings: newAddOnShipping,
		StatusShipping: "PROCESS",
	}

	shippingResposne := dto.CreateShippingResponse(shipping)

	mockAddOnRepo.On("GetAddOnByMultipleId", shippingCreateRequest.AddOnsId).Return(addOns, nil)
	mockSizeRepo.On("GetSizeById", shippingCreateRequest.SizeId).Return(&size, nil)
	mockCategoryRepo.On("GetCategoryById", shippingCreateRequest.CategoryId).Return(&category, nil)
	mockUserRepo.On("GetUserById", shippingCreateRequest.UserId).Return(&user, nil)
	mockAddressRepo.On("GetAddressByUserId", shippingCreateRequest.UserId, shippingCreateRequest.AddressId).Return(&address, nil)
	mockPaymentRepo.On("CreatePayment", payment).Return(&payment, nil)
	mockShippingRepo.On("CreateShipping", shipping).Return(&shipping, nil)

	res, err := usecase.CreateShipping(shippingCreateRequest)

	assert.Nil(t, err)
	assert.Equal(t, &shippingResposne, res)
}
