package usecase

import "final-project-backend/dto"

type CategoryUsecase interface {
	GetAllCategory() ([]dto.CategoryResponse, error)
}
