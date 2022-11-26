package repository

import "final-project-backend/entity"

type PromoUserRepository interface {
	GetPromoUserById(promoUserId int) (*entity.PromoUser, error)
	UpdatePromoUser(promo entity.PromoUser) (*entity.PromoUser, error)
}
