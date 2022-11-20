package usecase

import "final-project-backend/dto"

type SizeUsecase interface {
	GetAllSize() ([]dto.SizeResponse, error)
}
