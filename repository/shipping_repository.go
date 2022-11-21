package repository

import "final-project-backend/entity"

type ShippingRepository interface {
	CreateShipping(shipping entity.Shipping) (*entity.Shipping, error)
}
