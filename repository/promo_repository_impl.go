package repository

import (
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"

	"gorm.io/gorm"
)

type PromoRepositoryImpl struct {
	db *gorm.DB
}

func NewPromoRepositoryImpl(db *gorm.DB) PromoRepository {
	return &PromoRepositoryImpl{
		db: db,
	}
}

func (r *PromoRepositoryImpl) GetPromoById(promoId int) (*entity.Promo, error) {
	var promo entity.Promo

	err := r.db.First(&promo, promoId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custErr.ErrPromoNotFound
		}

		return nil, err
	}

	return nil, nil
}
