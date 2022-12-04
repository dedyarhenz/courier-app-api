package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryUsecaseImpl_GetAllCategory_Error(t *testing.T) {
	mockRepoCategory := mocks.NewCategoryRepository(t)
	usecase := NewCategoryUsecaseImpl(mockRepoCategory)

	mockRepoCategory.On("GetAllCategory").Return(nil, fmt.Errorf("error"))

	res, err := usecase.GetAllCategory()

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestCategoryUsecaseImpl_GetAllCategory_Success(t *testing.T) {
	mockRepoCategory := mocks.NewCategoryRepository(t)
	usecase := NewCategoryUsecaseImpl(mockRepoCategory)

	categories := []entity.Category{
		{Id: 1, Name: "category 1", Description: "desc category", Price: 10000},
		{Id: 2, Name: "category 2", Description: "desc category", Price: 20000},
	}
	categoriesDto := dto.CreateCategoryListResponse(categories)

	mockRepoCategory.On("GetAllCategory").Return(categories, nil)

	res, err := usecase.GetAllCategory()

	assert.Nil(t, err)
	assert.Equal(t, categoriesDto, res)
}
