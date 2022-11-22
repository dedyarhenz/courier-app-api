package repository

import "final-project-backend/entity"

type ShippingRepository interface {
	GetAllShippingByUserId(userId int) ([]entity.Shipping, error)
	GetShippingByUserId(userId int, shippingId int) (*entity.Shipping, error)
	CreateShipping(shipping entity.Shipping) (*entity.Shipping, error)
	UpdateReviewByUserId(userId int, shippingId int, review string) error
}
