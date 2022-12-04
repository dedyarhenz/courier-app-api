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

func TestPaymentUsecaseImpl_PayUserShipping_ErrorPaymentId(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", 1).Return(nil, fmt.Errorf("error"))
	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorPaymentNotFound(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 2,
			},
		},
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", 1).Return(&payment, nil)
	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, custErr.ErrPaymentNotFound, err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorAlreadyPaid(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "SUCCESS",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", 1).Return(&payment, nil)
	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, custErr.ErrShippingAlreadyPaid, err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorUserNotFound(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", 1).Return(&payment, nil)
	mockUserRepo.On("GetUserById", 1).Return(nil, fmt.Errorf("error"))
	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorPromoNotFound(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	user := entity.User{
		Id:      1,
		Email:   "dedy@gmail.com",
		Role:    "USER",
		Balance: 0,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(nil, fmt.Errorf("error"))

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorPromoUserNotFound(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	user := entity.User{
		Id:      1,
		Email:   "dedy@gmail.com",
		Role:    "USER",
		Balance: 0,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	promoUser := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  2,
		IsUsed:  false,
		Promo: &entity.Promo{
			ExpireDate: time.Date(2022, 12, 10, 9, 9, 9, 9, &time.Location{}),
			MinFee:     100000,
		},
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(&promoUser, nil)

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, custErr.ErrPromoNotFound, err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorPromoUsed(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	user := entity.User{
		Id:      1,
		Email:   "dedy@gmail.com",
		Role:    "USER",
		Balance: 0,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	promoUser := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  paymentRequest.UserId,
		IsUsed:  true,
		Promo: &entity.Promo{
			ExpireDate: time.Date(2022, 12, 10, 9, 9, 9, 9, &time.Location{}),
			MinFee:     100000,
		},
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(&promoUser, nil)

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, custErr.ErrPromoAlreadyUsed, err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorPromoExpired(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	user := entity.User{
		Id:      1,
		Email:   "dedy@gmail.com",
		Role:    "USER",
		Balance: 0,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	promoUser := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  paymentRequest.UserId,
		IsUsed:  false,
		Promo: &entity.Promo{
			ExpireDate: time.Date(2022, 11, 10, 9, 9, 9, 9, &time.Location{}),
			MinFee:     100000,
		},
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(&promoUser, nil)

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, custErr.ErrQuotaExpired, err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorPromoMinFee(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	user := entity.User{
		Id:      1,
		Email:   "dedy@gmail.com",
		Role:    "USER",
		Balance: 0,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	promoUser := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  paymentRequest.UserId,
		IsUsed:  false,
		Promo: &entity.Promo{
			ExpireDate: time.Date(2022, 12, 10, 9, 9, 9, 9, &time.Location{}),
			MinFee:     200000,
		},
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(&promoUser, nil)

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, custErr.ErrPromoFeeInvalid, err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorPromoUpdateUsed(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	user := entity.User{
		Id:      1,
		Email:   "dedy@gmail.com",
		Role:    "USER",
		Balance: 0,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	promoUser := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  paymentRequest.UserId,
		IsUsed:  false,
		Promo: &entity.Promo{
			Discount:    50,
			MaxDiscount: 10000,
			ExpireDate:  time.Date(2022, 12, 10, 9, 9, 9, 9, &time.Location{}),
			MinFee:      100000,
		},
	}

	promoUserUpdated := promoUser
	promoUserUpdated.IsUsed = true

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(&promoUser, nil)
	mockPromoUserRepo.On("UpdatePromoUser", promoUserUpdated).Return(nil, fmt.Errorf("error"))

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorBalance(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	user := entity.User{
		Id:      1,
		Email:   "dedy@gmail.com",
		Role:    "USER",
		Balance: 0,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	promoUser := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  paymentRequest.UserId,
		IsUsed:  false,
		Promo: &entity.Promo{
			Discount:    50,
			MaxDiscount: 10000,
			ExpireDate:  time.Date(2022, 12, 10, 9, 9, 9, 9, &time.Location{}),
			MinFee:      100000,
		},
	}

	promoUserUpdated := promoUser
	promoUserUpdated.IsUsed = true

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(&promoUser, nil)
	mockPromoUserRepo.On("UpdatePromoUser", promoUserUpdated).Return(&promoUserUpdated, nil)

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, custErr.ErrInsufficientBalance, err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorUpdatePayment(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	user := entity.User{
		Id:      1,
		Email:   "dedy@gmail.com",
		Role:    "USER",
		Balance: 500000,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	promoUser := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  paymentRequest.UserId,
		IsUsed:  false,
		Promo: &entity.Promo{
			Discount:    50,
			MaxDiscount: 10000,
			ExpireDate:  time.Date(2022, 12, 10, 9, 9, 9, 9, &time.Location{}),
			MinFee:      100000,
		},
	}

	promoUserUpdated := promoUser
	promoUserUpdated.IsUsed = true

	newPayment := entity.Payment{
		Id:            paymentRequest.PaymentId,
		PaymentStatus: "SUCCESS",
		TotalCost:     90000,
		PromoId:       &promoUser.PromoId,
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(&promoUser, nil)
	mockPromoUserRepo.On("UpdatePromoUser", promoUserUpdated).Return(&promoUserUpdated, nil)
	mockPaymentRepo.On("UpdatePayment", newPayment).Return(nil, fmt.Errorf("error"))

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorReduceBalance(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	user := entity.User{
		Id:      1,
		Email:   "dedy@gmail.com",
		Role:    "USER",
		Balance: 500000,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	promoUser := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  paymentRequest.UserId,
		IsUsed:  false,
		Promo: &entity.Promo{
			Discount:    50,
			MaxDiscount: 10000,
			ExpireDate:  time.Date(2022, 12, 10, 9, 9, 9, 9, &time.Location{}),
			MinFee:      100000,
		},
	}

	promoUserUpdated := promoUser
	promoUserUpdated.IsUsed = true

	newPayment := entity.Payment{
		Id:            paymentRequest.PaymentId,
		PaymentStatus: "SUCCESS",
		TotalCost:     90000,
		PromoId:       &promoUser.PromoId,
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(&promoUser, nil)
	mockPromoUserRepo.On("UpdatePromoUser", promoUserUpdated).Return(&promoUserUpdated, nil)
	mockPaymentRepo.On("UpdatePayment", newPayment).Return(&newPayment, nil)
	mockUserRepo.On("ReduceBalance", paymentRequest.UserId, newPayment.TotalCost).Return(nil, fmt.Errorf("error"))

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorBonusAddBalance(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	reffUserId := 2
	user := entity.User{
		Id:                  1,
		Email:               "dedy@gmail.com",
		Role:                "USER",
		Balance:             500000,
		RefferedUserId:      &reffUserId,
		IsBonusComplete:     false,
		IsBonusCompleteReff: false,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	promoUser := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  paymentRequest.UserId,
		IsUsed:  false,
		Promo: &entity.Promo{
			Discount:    50,
			MaxDiscount: 10000,
			ExpireDate:  time.Date(2022, 12, 10, 9, 9, 9, 9, &time.Location{}),
			MinFee:      100000,
		},
	}

	promoUserUpdated := promoUser
	promoUserUpdated.IsUsed = true

	newPayment := entity.Payment{
		Id:            paymentRequest.PaymentId,
		PaymentStatus: "SUCCESS",
		TotalCost:     90000,
		PromoId:       &promoUser.PromoId,
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(&promoUser, nil)
	mockPromoUserRepo.On("UpdatePromoUser", promoUserUpdated).Return(&promoUserUpdated, nil)
	mockPaymentRepo.On("UpdatePayment", newPayment).Return(&newPayment, nil)
	mockUserRepo.On("ReduceBalance", paymentRequest.UserId, newPayment.TotalCost).Return(&user, nil)
	mockPaymentRepo.On("TotalCostPaymentSuccessUser", paymentRequest.UserId).Return(int64(1000000))
	mockUserRepo.On("AddBalance", paymentRequest.UserId, 50000).Return(nil, fmt.Errorf("error"))

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorUpdateCompleteBonus(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	reffUserId := 2
	user := entity.User{
		Id:                  1,
		Email:               "dedy@gmail.com",
		Role:                "USER",
		Balance:             500000,
		RefferedUserId:      &reffUserId,
		IsBonusComplete:     false,
		IsBonusCompleteReff: false,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	promoUser := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  paymentRequest.UserId,
		IsUsed:  false,
		Promo: &entity.Promo{
			Discount:    50,
			MaxDiscount: 10000,
			ExpireDate:  time.Date(2022, 12, 10, 9, 9, 9, 9, &time.Location{}),
			MinFee:      100000,
		},
	}

	promoUserUpdated := promoUser
	promoUserUpdated.IsUsed = true

	newPayment := entity.Payment{
		Id:            paymentRequest.PaymentId,
		PaymentStatus: "SUCCESS",
		TotalCost:     90000,
		PromoId:       &promoUser.PromoId,
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(&promoUser, nil)
	mockPromoUserRepo.On("UpdatePromoUser", promoUserUpdated).Return(&promoUserUpdated, nil)
	mockPaymentRepo.On("UpdatePayment", newPayment).Return(&newPayment, nil)
	mockUserRepo.On("ReduceBalance", paymentRequest.UserId, newPayment.TotalCost).Return(&user, nil)
	mockPaymentRepo.On("TotalCostPaymentSuccessUser", paymentRequest.UserId).Return(int64(1000000))
	mockUserRepo.On("AddBalance", paymentRequest.UserId, 50000).Return(&user, nil)
	mockUserRepo.On("UpdatedCompleteBonus", paymentRequest.UserId, true).Return(fmt.Errorf("error"))

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorUserReff(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	reffUserId := 2
	user := entity.User{
		Id:                  1,
		Email:               "dedy@gmail.com",
		Role:                "USER",
		Balance:             500000,
		RefferedUserId:      &reffUserId,
		IsBonusComplete:     false,
		IsBonusCompleteReff: false,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	promoUser := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  paymentRequest.UserId,
		IsUsed:  false,
		Promo: &entity.Promo{
			Discount:    50,
			MaxDiscount: 10000,
			ExpireDate:  time.Date(2022, 12, 10, 9, 9, 9, 9, &time.Location{}),
			MinFee:      100000,
		},
	}

	promoUserUpdated := promoUser
	promoUserUpdated.IsUsed = true

	newPayment := entity.Payment{
		Id:            paymentRequest.PaymentId,
		PaymentStatus: "SUCCESS",
		TotalCost:     90000,
		PromoId:       &promoUser.PromoId,
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(&promoUser, nil)
	mockPromoUserRepo.On("UpdatePromoUser", promoUserUpdated).Return(&promoUserUpdated, nil)
	mockPaymentRepo.On("UpdatePayment", newPayment).Return(&newPayment, nil)
	mockUserRepo.On("ReduceBalance", paymentRequest.UserId, newPayment.TotalCost).Return(&user, nil)
	mockPaymentRepo.On("TotalCostPaymentSuccessUser", paymentRequest.UserId).Return(int64(1000000))
	mockUserRepo.On("AddBalance", paymentRequest.UserId, 50000).Return(&user, nil)
	mockUserRepo.On("UpdatedCompleteBonus", paymentRequest.UserId, true).Return(nil)
	mockUserRepo.On("GetUserById", *user.RefferedUserId).Return(nil, fmt.Errorf("error"))

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorUserReffBonusAddBalance(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	reffUserId := 2
	user := entity.User{
		Id:                  1,
		Email:               "dedy@gmail.com",
		Role:                "USER",
		Balance:             500000,
		RefferedUserId:      &reffUserId,
		IsBonusComplete:     false,
		IsBonusCompleteReff: false,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	promoUser := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  paymentRequest.UserId,
		IsUsed:  false,
		Promo: &entity.Promo{
			Discount:    50,
			MaxDiscount: 10000,
			ExpireDate:  time.Date(2022, 12, 10, 9, 9, 9, 9, &time.Location{}),
			MinFee:      100000,
		},
	}

	promoUserUpdated := promoUser
	promoUserUpdated.IsUsed = true

	newPayment := entity.Payment{
		Id:            paymentRequest.PaymentId,
		PaymentStatus: "SUCCESS",
		TotalCost:     90000,
		PromoId:       &promoUser.PromoId,
	}

	userReff := entity.User{
		Id:      *user.RefferedUserId,
		Balance: 500000,
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(&promoUser, nil)
	mockPromoUserRepo.On("UpdatePromoUser", promoUserUpdated).Return(&promoUserUpdated, nil)
	mockPaymentRepo.On("UpdatePayment", newPayment).Return(&newPayment, nil)
	mockUserRepo.On("ReduceBalance", paymentRequest.UserId, newPayment.TotalCost).Return(&user, nil)
	mockPaymentRepo.On("TotalCostPaymentSuccessUser", paymentRequest.UserId).Return(int64(1000000))
	mockUserRepo.On("AddBalance", paymentRequest.UserId, 50000).Return(&user, nil)
	mockUserRepo.On("UpdatedCompleteBonus", paymentRequest.UserId, true).Return(nil)
	mockUserRepo.On("GetUserById", *user.RefferedUserId).Return(&userReff, nil)
	mockUserRepo.On("AddBalance", *user.RefferedUserId, 25000).Return(nil, fmt.Errorf("error"))

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPaymentUsecaseImpl_PayUserShipping_ErrorUserReffUpdateCompleteBonus(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	reffUserId := 2
	user := entity.User{
		Id:                  1,
		Email:               "dedy@gmail.com",
		Role:                "USER",
		Balance:             500000,
		RefferedUserId:      &reffUserId,
		IsBonusComplete:     false,
		IsBonusCompleteReff: false,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	promoUser := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  paymentRequest.UserId,
		IsUsed:  false,
		Promo: &entity.Promo{
			Discount:    50,
			MaxDiscount: 10000,
			ExpireDate:  time.Date(2022, 12, 10, 9, 9, 9, 9, &time.Location{}),
			MinFee:      100000,
		},
	}

	promoUserUpdated := promoUser
	promoUserUpdated.IsUsed = true

	newPayment := entity.Payment{
		Id:            paymentRequest.PaymentId,
		PaymentStatus: "SUCCESS",
		TotalCost:     90000,
		PromoId:       &promoUser.PromoId,
	}

	userReff := entity.User{
		Id:      *user.RefferedUserId,
		Balance: 500000,
	}

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(&promoUser, nil)
	mockPromoUserRepo.On("UpdatePromoUser", promoUserUpdated).Return(&promoUserUpdated, nil)
	mockPaymentRepo.On("UpdatePayment", newPayment).Return(&newPayment, nil)
	mockUserRepo.On("ReduceBalance", paymentRequest.UserId, newPayment.TotalCost).Return(&user, nil)
	mockPaymentRepo.On("TotalCostPaymentSuccessUser", paymentRequest.UserId).Return(int64(1000000))
	mockUserRepo.On("AddBalance", paymentRequest.UserId, 50000).Return(&user, nil)
	mockUserRepo.On("UpdatedCompleteBonus", paymentRequest.UserId, true).Return(nil)
	mockUserRepo.On("GetUserById", *user.RefferedUserId).Return(&userReff, nil)
	mockUserRepo.On("AddBalance", *user.RefferedUserId, 25000).Return(&userReff, nil)
	mockUserRepo.On("UpdatedCompleteBonusReff", *user.RefferedUserId, true).Return(fmt.Errorf("error"))

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPaymentUsecaseImpl_PayUserShipping_Success(t *testing.T) {
	mockPaymentRepo := mocks.NewPaymentRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)

	var promoUserId int = 1
	paymentRequest := dto.PaymentPayRequest{
		PaymentId:   1,
		UserId:      1,
		PromoUserId: &promoUserId,
	}

	reffUserId := 2
	user := entity.User{
		Id:                  1,
		Email:               "dedy@gmail.com",
		Role:                "USER",
		Balance:             500000,
		RefferedUserId:      &reffUserId,
		IsBonusComplete:     false,
		IsBonusCompleteReff: false,
	}

	payment := entity.Payment{
		Id:            1,
		PaymentStatus: "PENDING",
		TotalCost:     100000,
		PromoId:       &paymentRequest.PaymentId,
		Promo:         nil,
		Shipping: &entity.Shipping{
			Address: &entity.Address{
				UserId: 1,
			},
		},
	}

	promoUser := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  paymentRequest.UserId,
		IsUsed:  false,
		Promo: &entity.Promo{
			Discount:    50,
			MaxDiscount: 10000,
			ExpireDate:  time.Date(2022, 12, 10, 9, 9, 9, 9, &time.Location{}),
			MinFee:      100000,
		},
	}

	promoUserUpdated := promoUser
	promoUserUpdated.IsUsed = true

	newPayment := entity.Payment{
		Id:            paymentRequest.PaymentId,
		PaymentStatus: "SUCCESS",
		TotalCost:     90000,
		PromoId:       &promoUser.PromoId,
	}

	userReff := entity.User{
		Id:      *user.RefferedUserId,
		Balance: 500000,
	}

	paymentResponse := dto.CreatePaymentResponse(newPayment)

	usecase := NewPaymentUsecaseImpl(mockPaymentRepo, mockUserRepo, mockPromoUserRepo)

	mockPaymentRepo.On("GetPaymentById", paymentRequest.PaymentId).Return(&payment, nil)
	mockUserRepo.On("GetUserById", paymentRequest.UserId).Return(&user, nil)
	mockPromoUserRepo.On("GetPromoUserById", *paymentRequest.PromoUserId).Return(&promoUser, nil)
	mockPromoUserRepo.On("UpdatePromoUser", promoUserUpdated).Return(&promoUserUpdated, nil)
	mockPaymentRepo.On("UpdatePayment", newPayment).Return(&newPayment, nil)
	mockUserRepo.On("ReduceBalance", paymentRequest.UserId, newPayment.TotalCost).Return(&user, nil)
	mockPaymentRepo.On("TotalCostPaymentSuccessUser", paymentRequest.UserId).Return(int64(1000000))
	mockUserRepo.On("AddBalance", paymentRequest.UserId, 50000).Return(&user, nil)
	mockUserRepo.On("UpdatedCompleteBonus", paymentRequest.UserId, true).Return(nil)
	mockUserRepo.On("GetUserById", *user.RefferedUserId).Return(&userReff, nil)
	mockUserRepo.On("AddBalance", *user.RefferedUserId, 25000).Return(&userReff, nil)
	mockUserRepo.On("UpdatedCompleteBonusReff", *user.RefferedUserId, true).Return(nil)

	res, err := usecase.PayUserShipping(paymentRequest)

	assert.Nil(t, err)
	assert.Equal(t, &paymentResponse, res)
}
