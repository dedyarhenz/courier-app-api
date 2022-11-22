package repository

import (
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	err := r.db.Clauses(clause.Returning{}).Omit("created_at", "updated_at", "deleted_at").Create(&newShipping).Error
	if err != nil {
		return nil, err
	}

	return &newShipping, nil
}

func (r *ShippingRepositoryImpl) GetAllShippingByUserId(userId int) ([]entity.Shipping, error) {
	var shippings []entity.Shipping

	err := r.db.
		Joins("INNER JOIN addresses ON addresses.id = shippings.address_id AND addresses.user_id = ?", userId).
		Preload("Address").
		Preload("Size").
		Preload("Category").
		Preload("Payment").
		Find(&shippings).Error

	if err != nil {
		return nil, err
	}

	return shippings, nil
}

func (r *ShippingRepositoryImpl) GetShippingByUserId(userId int, shippingId int) (*entity.Shipping, error) {
	var shipping entity.Shipping

	err := r.db.
		Joins("INNER JOIN addresses ON addresses.id = shippings.address_id AND addresses.user_id = ?", userId).
		Preload("Address").
		Preload("Size").
		Preload("Category").
		Preload("Payment").
		Preload("AddOns").
		Where("shippings.id", shippingId).
		First(&shipping).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custErr.ErrShippingNotFound
		}

		return nil, err
	}

	return &shipping, nil
}

func (r *ShippingRepositoryImpl) UpdateReviewByUserId(userId int, shippingId int, review string) error {
	res := r.db.
		Model(&entity.Shipping{}).
		Joins("INNER JOIN addresses ON addresses.id = shippings.address_id AND shippings.user_id = ?", userId).
		Where("shippings.id", shippingId).
		UpdateColumn("review", review)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return custErr.ErrShippingNotFound
	}

	return nil
}
