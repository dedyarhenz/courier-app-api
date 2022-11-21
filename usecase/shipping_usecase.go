package usecase

import (
	"final-project-backend/dto"
)

type ShippingUsecase interface {
	CreateShipping(request dto.ShippingCreateRequest) (*dto.ShippingResponse, error)
}
