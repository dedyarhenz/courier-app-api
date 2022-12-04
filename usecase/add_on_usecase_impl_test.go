package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOnUsecaseImpl_GetAllAddOn_Error(t *testing.T) {
	mockRepoAddOn := mocks.NewAddOnRepository(t)
	usecase := NewAddOnUsecaseImpl(mockRepoAddOn)

	mockRepoAddOn.On("GetAllAddOn").Return(nil, fmt.Errorf("error"))

	res, err := usecase.GetAllAddOn()

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestAddOnUsecaseImpl_GetAllAddOn_Success(t *testing.T) {
	mockRepoAddOn := mocks.NewAddOnRepository(t)
	usecase := NewAddOnUsecaseImpl(mockRepoAddOn)

	addOns := []entity.AddOn{
		{Id: 1, Name: "addon 1", Description: "desc addon", Price: 10000},
		{Id: 2, Name: "addon 2", Description: "desc addon", Price: 20000},
	}
	addOnsDto := dto.CreateAddOnListResponse(addOns)

	mockRepoAddOn.On("GetAllAddOn").Return(addOns, nil)

	res, err := usecase.GetAllAddOn()

	assert.Nil(t, err)
	assert.Equal(t, addOnsDto, res)
}
