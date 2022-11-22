package repository

import "final-project-backend/entity"

type ShippingRepository interface {
	GetShippingByUserId(userId int) ([]entity.Shipping, error)
	CreateShipping(shipping entity.Shipping) (*entity.Shipping, error)
}
