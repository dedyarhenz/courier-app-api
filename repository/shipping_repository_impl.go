package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type ShippingRepositoryImpl struct {
	db *gorm.DB
}

func NewShippingRepositoryImpl(db *gorm.DB) ShippingRepository {
	return &ShippingRepositoryImpl{
		db: db,
	}
}

func (r *ShippingRepositoryImpl) CreateShipping(shipping entity.Shipping) (*entity.Shipping, error) {
	newShipping := entity.Shipping{
		SizeId:         shipping.SizeId,
		CategoryId:     shipping.CategoryId,
		AddressId:      shipping.AddressId,
		PaymentId:      shipping.PaymentId,
		StatusShipping: shipping.StatusShipping,
		Review:         shipping.Review,
		AddOnShippings: shipping.AddOnShippings,
	}

	err := r.db.Omit("created_at", "updated_at", "deleted_at").Create(&newShipping).Error
	if err != nil {
		return nil, err
	}

	return &newShipping, nil
}
