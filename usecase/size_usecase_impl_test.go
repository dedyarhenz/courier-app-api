package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSizeUsecaseImpl_GetAllSize_Error(t *testing.T) {
	mockRepoSize := mocks.NewSizeRepository(t)
	usecase := NewSizeUsecaseImpl(mockRepoSize)

	mockRepoSize.On("GetAllSize").Return(nil, fmt.Errorf("error"))

	res, err := usecase.GetAllSize()

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestSizeUsecaseImpl_GetAllSize_Success(t *testing.T) {
	mockRepoSize := mocks.NewSizeRepository(t)
	usecase := NewSizeUsecaseImpl(mockRepoSize)

	sizes := []entity.Size{
		{Id: 1, Name: "size 1", Description: "desc size", Price: 10000},
		{Id: 2, Name: "size 2", Description: "desc size", Price: 20000},
	}
	sizesDto := dto.CreateSizeListResponse(sizes)

	mockRepoSize.On("GetAllSize").Return(sizes, nil)

	res, err := usecase.GetAllSize()

	assert.Nil(t, err)
	assert.Equal(t, sizesDto, res)
}
