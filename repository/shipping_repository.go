package repository

import (
	"final-project-backend/entity"
)

type ShippingRepository interface {
	GetAllShipping(offset int, limit int, search string, orderAndSort string) ([]entity.Shipping, error)
	GetAllReportShippingByDate(startDate string, endDate string, offset int, limit int, orderAndSort string) ([]entity.Shipping, error)
	GetShippingById(shippingId int) (*entity.Shipping, error)
	GetAllShippingByUserId(userId int, offset int, limit int, search string, orderAndSort string) ([]entity.Shipping, error)
	GetShippingByUserId(userId int, shippingId int) (*entity.Shipping, error)
	CreateShipping(shipping entity.Shipping) (*entity.Shipping, error)
	UpdateShipping(shipping entity.Shipping) (*entity.Shipping, error)
	UpdateReviewByUserId(userId int, shippingId int, review string) error
	UpdateStatusShipping(shippingId int, statusShipping string) error
	CountShipping(search string) int64
	CountShippingByUserId(userId int, search string) int64
	CountShippingByDate(startDate string, endDate string) int64
}
