package usecase

import (
	"final-project-backend/dto"
)

type ShippingUsecase interface {
	CreateShipping(request dto.ShippingCreateRequest) (*dto.ShippingResponse, error)
	GetAllShipping(page int, limit int, search string, order string, sortBy string) (dto.ShippingPaginateResponse, error)
	GetAllReportShippingByDate(month int, year int, page int, limit int, sortBy string) (dto.ShippingReportPaginateResponse, error)
	GetShippingById(shippingId int) (*dto.ShippingResponse, error)
	GetAllShippingByUserId(userId int, page int, limit int, search string, order string, sortBy string) (dto.ShippingPaginateResponse, error)
	GetShippingByUserId(userId int, shippingId int) (*dto.ShippingResponse, error)
	UpdateReviewByUserId(request dto.ShippingReviewRequest) error
	UpdateStatusShipping(request dto.ShippingUpdateStatusRequest) error
}
