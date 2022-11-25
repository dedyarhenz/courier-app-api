package repository

import "final-project-backend/entity"

type PromoRepository interface {
	GetAllPromo(offset int, limit int, search string, orderAndSort string) ([]entity.Promo, error)
	GetPromoById(promoId int) (*entity.Promo, error)
	CreatePromo(promo entity.Promo) (*entity.Promo, error)
	UpdatePromo(promo entity.Promo) (*entity.Promo, error)
	CountPromo(search string) int64
}
