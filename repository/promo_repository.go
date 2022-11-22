package repository

import "final-project-backend/entity"

type PromoRepository interface {
	GetPromoById(promoId int) (*entity.Promo, error)
}
