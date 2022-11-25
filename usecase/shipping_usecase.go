package usecase

import (
	"final-project-backend/dto"
)

type ShippingUsecase interface {
	CreateShipping(request dto.ShippingCreateRequest) (*dto.ShippingResponse, error)
	GetAllShipping(page int, limit int, search string, order string, sortBy string) (dto.ShippingPaginateResponse, error)
	GetAllShippingByUserId(userId int, page int, limit int, search string, order string, sortBy string) (dto.ShippingPaginateResponse, error)
	GetShippingByUserId(userId int, shippingId int) (*dto.ShippingResponse, error)
	UpdateReviewByUserId(request dto.ShippingReviewRequest) error
}
