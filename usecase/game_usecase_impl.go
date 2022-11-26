package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"
	"final-project-backend/repository"
	"math/rand"
	"time"
)

type GameUsecaseImpl struct {
	repoPromo     repository.PromoRepository
	repoPromoUser repository.PromoUserRepository
	repoShipping  repository.ShippingRepository
}

func NewGameUsecaseImpl(
	repoPromo repository.PromoRepository,
	repoPromoUser repository.PromoUserRepository,
	repoShipping repository.ShippingRepository,
) GameUsecase {
	return &GameUsecaseImpl{
		repoPromo:     repoPromo,
		repoPromoUser: repoPromoUser,
		repoShipping:  repoShipping,
	}
}

func (u *GameUsecaseImpl) Play(request dto.GamePlayRequest) (*dto.GameResponse, error) {
	promos, err := u.repoPromo.GetAllPromoGame()
	if err != nil {
		return nil, err
	}

	shipping, err := u.repoShipping.GetShippingByUserId(request.UserId, request.ShippingId)
	if err != nil {
		return nil, err
	}

	if shipping.StatusShipping != entity.SHIPP_DELIVERED || shipping.Payment.PaymentStatus != entity.PAYMENT_SUCCESS {
		return nil, custErr.ErrGameTransactionNotDone
	}

	if shipping.Payment.TotalCost < 20000 {
		return nil, custErr.ErrGameMinTransaction
	}

	if shipping.IsPlayGame {
		return nil, custErr.ErrGameChanceUsed
	}

	rand.Seed(time.Now().Unix())
	promo := promos[rand.Intn(len(promos))]

	newPromoUser := entity.PromoUser{
		PromoId: promo.Id,
		UserId:  request.UserId,
	}

	_, err = u.repoPromoUser.CreatePromoUser(newPromoUser)
	if err != nil {
		return nil, err
	}

	promo.Quota = promo.Quota - 1
	_, err = u.repoPromo.UpdatePromo(promo)
	if err != nil {
		return nil, err
	}

	shipping.IsPlayGame = true
	_, err = u.repoShipping.UpdateShipping(*shipping)
	if err != nil {
		return nil, err
	}

	resGame := dto.GameResponse{
		Name:        promo.Name,
		MinFee:      promo.MinFee,
		Discount:    promo.Discount,
		MaxDiscount: promo.MaxDiscount,
		ExpireDate:  promo.ExpireDate,
	}

	return &resGame, nil
}
