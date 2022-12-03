package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"
	"final-project-backend/repository"
	"time"
)

type PaymentUsecaseImpl struct {
	repoPayment   repository.PaymentRepository
	repoUser      repository.UserRepository
	repoPromoUser repository.PromoUserRepository
}

func NewPaymentUsecaseImpl(
	repoPayment repository.PaymentRepository,
	repoUser repository.UserRepository,
	repoPromoUser repository.PromoUserRepository,
) PaymenUsecase {
	return &PaymentUsecaseImpl{
		repoPayment:   repoPayment,
		repoUser:      repoUser,
		repoPromoUser: repoPromoUser,
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

	user, err := u.repoUser.GetUserById(request.UserId)
	if err != nil {
		return nil, err
	}

	var totalDiscount int = 0
	var promoId *int
	if request.PromoUserId != nil {
		promoUser, err := u.repoPromoUser.GetPromoUserById(*request.PromoUserId)
		if err != nil {
			return nil, err
		}

		totalDiscount, err = u.caclulatePromo(*promoUser, request.UserId, payment.TotalCost)
		if err != nil {
			return nil, err
		}

		promoUser.IsUsed = true
		_, err = u.repoPromoUser.UpdatePromoUser(*promoUser)
		if err != nil {
			return nil, err
		}

		promoId = &promoUser.PromoId
	}

	totalPrice := payment.TotalCost - totalDiscount

	if user.Balance < totalPrice {
		return nil, custErr.ErrInsufficientBalance
	}

	newPayment := entity.Payment{
		Id:            request.PaymentId,
		PaymentStatus: entity.PAYMENT_SUCCESS,
		TotalCost:     totalPrice,
		PromoId:       promoId,
	}
	paymentUpdate, err := u.repoPayment.UpdatePayment(newPayment)
	if err != nil {
		return nil, err
	}

	_, err = u.repoUser.ReduceBalance(request.UserId, totalPrice)
	if err != nil {
		return nil, err
	}

	totalCostPaymentSuccess := u.repoPayment.TotalCostPaymentSuccessUser(request.UserId)

	if totalCostPaymentSuccess >= 350000 && user.RefferedUserId != nil && !user.IsBonusComplete {
		_, err = u.repoUser.AddBalance(request.UserId, 50000)
		if err != nil {
			return nil, err
		}

		err = u.repoUser.UpdatedCompleteBonus(request.UserId, true)
		if err != nil {
			return nil, err
		}
	}

	if totalCostPaymentSuccess >= 500000 && user.RefferedUserId != nil {
		userReff, err := u.repoUser.GetUserById(*user.RefferedUserId)
		if err != nil {
			return nil, err
		}

		if !userReff.IsBonusCompleteReff {
			_, err = u.repoUser.AddBalance(userReff.Id, 25000)
			if err != nil {
				return nil, err
			}

			err = u.repoUser.UpdatedCompleteBonusReff(userReff.Id, true)
			if err != nil {
				return nil, err
			}
		}

	}

	resPaymentUpdate := dto.CreatePaymentResponse(*paymentUpdate)

	return &resPaymentUpdate, nil
}

func (u *PaymentUsecaseImpl) caclulatePromo(promoUser entity.PromoUser, userId int, totalCost int) (int, error) {
	if promoUser.UserId != userId {
		return 0, custErr.ErrPromoNotFound
	}

	if promoUser.IsUsed {
		return 0, custErr.ErrPromoAlreadyUsed
	}

	now := time.Now()
	if now.After(promoUser.Promo.ExpireDate) {
		return 0, custErr.ErrQuotaExpired
	}

	if totalCost < promoUser.Promo.MinFee {
		return 0, custErr.ErrPromoFeeInvalid
	}

	totalDiscount := float64(totalCost) * (float64(promoUser.Promo.Discount) / float64(100))

	if totalDiscount > float64(promoUser.Promo.MaxDiscount) {
		totalDiscount = float64(promoUser.Promo.MaxDiscount)
	}

	return int(totalDiscount), nil
}
