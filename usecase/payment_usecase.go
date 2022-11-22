package usecase

import (
	"final-project-backend/dto"
)

type PaymenUsecase interface {
	PayUserShipping(request dto.PaymentPayRequest) (*dto.PaymentResponse, error)
}
