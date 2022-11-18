package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddressUsecaseImpl_GetAllAddress_Error(t *testing.T) {
	mockAddressRepo := mocks.NewAddressRepository(t)
	usecase := NewAddressUsecaseImpl(mockAddressRepo)

	mockAddressRepo.On("GetAllAddress").Return(nil, fmt.Errorf("error"))
	res, err := usecase.GetAllAddress()

	assert.Nil(t, res)
	assert.Error(t, fmt.Errorf("error"), err)
}

func TestAddressUsecaseImpl_GetAllAddress_Success(t *testing.T) {
	mockAddressRepo := mocks.NewAddressRepository(t)
	usecase := NewAddressUsecaseImpl(mockAddressRepo)

	addresses := []entity.Address{
		{
			Id:             1,
			RecipientName:  "budi",
			FullAddress:    "jl.qwerty",
			RecipientPhone: "zuzi",
			UserId:         2,
		},
		{
			Id:             2,
			RecipientName:  "budi2",
			FullAddress:    "jl.qwerty2",
			RecipientPhone: "zuzi2",
			UserId:         2,
		},
	}
	addressesResponse := dto.CreateAddressListResponse(addresses)

	mockAddressRepo.On("GetAllAddress").Return(addresses, nil)
	res, err := usecase.GetAllAddress()

	assert.Nil(t, err)
	assert.Equal(t, addressesResponse, res)
}
