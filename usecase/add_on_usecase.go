package usecase

import "final-project-backend/dto"

type AddOnUsecase interface {
	GetAllAddOn() ([]dto.AddOnResponse, error)
}
