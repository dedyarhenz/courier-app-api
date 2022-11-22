package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"
	"final-project-backend/repository"
	"time"
)

type PaymentUsecaseImpl struct {
	repoPayment repository.PaymentRepository
	repoUser    repository.UserRepository
	repoPromo   repository.PromoRepository
}

func NewPaymentUsecaseImpl(
	repoPayment repository.PaymentRepository,
	repoUser repository.UserRepository,
	repoPromo repository.PromoRepository,
) PaymenUsecase {
	return &PaymentUsecaseImpl{
		repoPayment: repoPayment,
		repoUser:    repoUser,
		repoPromo:   repoPromo,
	}
}

func (u *PaymentUsecaseImpl) PayUserShipping(request dto.PaymentPayRequest) (*dto.PaymentResponse, error) {
	payment, err := u.repoPayment.GetPaymentById(request.PaymentId)
	if err != nil {
		return nil, err
	}

	if payment.Shipping.Address.UserId != request.UserId {
		return nil, custErr.ErrPaymentNotFound
	}

	if payment.PaymentStatus == entity.PAYMENT_SUCCESS {
		return nil, custErr.ErrShippingAlreadyPaid
	}

	var totalDiscount int = 0
	var promo *entity.Promo
	if request.PromoId != nil {
		promo, err = u.repoPromo.GetPromoById(*request.PromoId)
		if err != nil {
			return nil, err
		}

		if promo.Quota == 0 {
			return nil, custErr.ErrQuotaOutOfStock
		}

		now := time.Now()
		if now.After(promo.ExpireDate) {
			return nil, custErr.ErrQuotaExpired
		}

		totalDiscount = payment.TotalCost * (promo.Discount / 100)
	}

	totalPrice := payment.TotalCost - totalDiscount

	user, err := u.repoUser.GetUserById(request.UserId)
	if err != nil {
		return nil, err
	}

	if user.Balance < totalPrice {
		return nil, custErr.ErrInsufficientBalance
	}

	newPayment := entity.Payment{
		Id:            request.PaymentId,
		PaymentStatus: entity.PAYMENT_SUCCESS,
		TotalCost:     totalPrice,
		PromoId:       request.PromoId,
	}
	paymentUpdate, err := u.repoPayment.UpdatePayment(newPayment)
	if err != nil {
		return nil, err
	}

	_, err = u.repoUser.ReduceBalance(request.UserId, totalPrice)
	if err != nil {
		return nil, err
	}

	if totalPrice >= 350000 && user.RefferedUserId != nil {
		_, err = u.repoUser.AddBalance(request.UserId, 50000)
		if err != nil {
			return nil, err
		}
	}

	if totalPrice >= 500000 && user.RefferedUserId != nil {
		_, err = u.repoUser.AddBalance(*user.RefferedUserId, 25000)
		if err != nil {
			return nil, err
		}
	}

	resPaymentUpdate := dto.CreatePaymentResponse(*paymentUpdate)

	return &resPaymentUpdate, nil
}
