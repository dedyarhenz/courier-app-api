package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/repository"
)

type CategoryUsecaseImpl struct {
	repoCategory repository.CategoryRepository
}

func NewCategoryUsecaseImpl(repoCategory repository.CategoryRepository) CategoryUsecase {
	return &CategoryUsecaseImpl{
		repoCategory: repoCategory,
	}
}

func (u *CategoryUsecaseImpl) GetAllCategory() ([]dto.CategoryResponse, error) {
	allCategory, err := u.repoCategory.GetAllCategory()
	if err != nil {
		return nil, err
	}

	resAllSize := dto.CreateCategoryListResponse(allCategory)

	return resAllSize, nil
}
