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

	mockAddressRepo.On("CountAddress", "").Return(int64(1))
	mockAddressRepo.On("GetAllAddress", 0, 10, "", "created_at desc").Return(nil, fmt.Errorf("error"))
	res, err := usecase.GetAllAddress(1, 10, "")

	assert.NotNil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
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
			UserId:         3,
		},
		{
			Id:             2,
			RecipientName:  "budi2",
			FullAddress:    "jl.qwerty2",
			RecipientPhone: "zuzi2",
			UserId:         3,
		},
	}
	addressesResponse := dto.CreateAddressListResponse(addresses)
	addressPaginate := dto.AddressPaginateResponse{
		Page:      1,
		Limit:     10,
		Totaldata: 2,
		TotalPage: 1,
		Data:      addressesResponse,
	}

	mockAddressRepo.On("CountAddress", "").Return(int64(2))
	mockAddressRepo.On("GetAllAddress", 0, 10, "", "created_at desc").Return(addresses, nil)
	res, err := usecase.GetAllAddress(1, 10, "")

	assert.Nil(t, err)
	assert.Equal(t, addressPaginate, res)
}

func TestAddressUsecaseImpl_GetAllAddressByUserId_Error(t *testing.T) {
	mockAddressRepo := mocks.NewAddressRepository(t)
	usecase := NewAddressUsecaseImpl(mockAddressRepo)

	mockAddressRepo.On("CountAddressByUserId", 3, "").Return(int64(1))
	mockAddressRepo.On("GetAllAddressByUserId", 3, 0, 10, "", "created_at desc").Return(nil, fmt.Errorf("error"))
	res, err := usecase.GetAllAddressByUserId(3, 1, 10, "")

	assert.NotNil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestAddressUsecaseImpl_GetAllAddressByUserId_Success(t *testing.T) {
	mockAddressRepo := mocks.NewAddressRepository(t)
	usecase := NewAddressUsecaseImpl(mockAddressRepo)

	addresses := []entity.Address{
		{
			Id:             1,
			RecipientName:  "budi",
			FullAddress:    "jl.qwerty",
			RecipientPhone: "zuzi",
			UserId:         3,
		},
		{
			Id:             2,
			RecipientName:  "budi2",
			FullAddress:    "jl.qwerty2",
			RecipientPhone: "zuzi2",
			UserId:         3,
		},
	}
	addressesResponse := dto.CreateAddressListResponse(addresses)
	addressPaginate := dto.AddressPaginateResponse{
		Page:      1,
		Limit:     10,
		Totaldata: 2,
		TotalPage: 1,
		Data:      addressesResponse,
	}

	mockAddressRepo.On("CountAddressByUserId", 3, "").Return(int64(2))
	mockAddressRepo.On("GetAllAddressByUserId", 3, 0, 10, "", "created_at desc").Return(addresses, nil)
	res, err := usecase.GetAllAddressByUserId(3, 1, 10, "")

	assert.Nil(t, err)
	assert.Equal(t, addressPaginate, res)
}

func TestAddressUsecaseImpl_GetAddressByUserId_Error(t *testing.T) {
	mockAddressRepo := mocks.NewAddressRepository(t)
	usecase := NewAddressUsecaseImpl(mockAddressRepo)

	mockAddressRepo.On("GetAddressByUserId", 3, 1).Return(nil, fmt.Errorf("error"))
	res, err := usecase.GetAddressByUserId(3, 1)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestAddressUsecaseImpl_GetAddressByUserId_Success(t *testing.T) {
	mockAddressRepo := mocks.NewAddressRepository(t)
	usecase := NewAddressUsecaseImpl(mockAddressRepo)

	address := entity.Address{
		Id:             1,
		RecipientName:  "budi",
		FullAddress:    "jl.qwerty",
		RecipientPhone: "zuzi",
		UserId:         3,
	}
	addressResponse := dto.CreateAddressResponse(address)

	mockAddressRepo.On("GetAddressByUserId", 3, 1).Return(&address, nil)
	res, err := usecase.GetAddressByUserId(3, 1)

	assert.Nil(t, err)
	assert.Equal(t, &addressResponse, res)
}

func TestAddressUsecaseImpl_CreateAddress_Error(t *testing.T) {
	mockAddressRepo := mocks.NewAddressRepository(t)
	usecase := NewAddressUsecaseImpl(mockAddressRepo)

	addressRequest := dto.AddressCreateRequest{
		RecipientName:  "budi",
		FullAddress:    "jl.qwerty",
		RecipientPhone: "zuzi",
		UserId:         3,
	}
	address := entity.Address{
		RecipientName:  "budi",
		FullAddress:    "jl.qwerty",
		RecipientPhone: "zuzi",
		UserId:         3,
	}

	mockAddressRepo.On("CreateAddress", address).Return(nil, fmt.Errorf("error"))
	res, err := usecase.CreateAddress(addressRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestAddressUsecaseImpl_CreateAddress_Success(t *testing.T) {
	mockAddressRepo := mocks.NewAddressRepository(t)
	usecase := NewAddressUsecaseImpl(mockAddressRepo)

	addressRequest := dto.AddressCreateRequest{
		RecipientName:  "budi",
		FullAddress:    "jl.qwerty",
		RecipientPhone: "zuzi",
		UserId:         3,
	}
	address := entity.Address{
		RecipientName:  "budi",
		FullAddress:    "jl.qwerty",
		RecipientPhone: "zuzi",
		UserId:         3,
	}
	addressResponse := dto.CreateAddressResponse(address)

	mockAddressRepo.On("CreateAddress", address).Return(&address, nil)
	res, err := usecase.CreateAddress(addressRequest)

	assert.Nil(t, err)
	assert.Equal(t, &addressResponse, res)
}

func TestAddressUsecaseImpl_UpdateAddressByUserId_Error(t *testing.T) {
	mockAddressRepo := mocks.NewAddressRepository(t)
	usecase := NewAddressUsecaseImpl(mockAddressRepo)

	addressRequest := dto.AddressUpdateRequest{
		Id:             1,
		RecipientName:  "budi",
		FullAddress:    "jl.qwerty",
		RecipientPhone: "zuzi",
		UserId:         3,
	}
	address := entity.Address{
		Id:             1,
		RecipientName:  "budi",
		FullAddress:    "jl.qwerty",
		RecipientPhone: "zuzi",
		UserId:         3,
	}

	mockAddressRepo.On("UpdateAddressByUserId", address).Return(nil, fmt.Errorf("error"))
	res, err := usecase.UpdateAddressByUserId(addressRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestAddressUsecaseImpl_UpdateAddressByUserId_Success(t *testing.T) {
	mockAddressRepo := mocks.NewAddressRepository(t)
	usecase := NewAddressUsecaseImpl(mockAddressRepo)

	addressRequest := dto.AddressUpdateRequest{
		Id:             1,
		RecipientName:  "budi",
		FullAddress:    "jl.qwerty",
		RecipientPhone: "zuzi",
		UserId:         3,
	}
	address := entity.Address{
		Id:             1,
		RecipientName:  "budi",
		FullAddress:    "jl.qwerty",
		RecipientPhone: "zuzi",
		UserId:         3,
	}
	addressResponse := dto.CreateAddressResponse(address)

	mockAddressRepo.On("UpdateAddressByUserId", address).Return(&address, nil)
	res, err := usecase.UpdateAddressByUserId(addressRequest)

	assert.Nil(t, err)
	assert.Equal(t, &addressResponse, res)
}

func TestAddressUsecaseImpl_DeleteAddressByUserId_Error(t *testing.T) {
	mockAddressRepo := mocks.NewAddressRepository(t)
	usecase := NewAddressUsecaseImpl(mockAddressRepo)

	mockAddressRepo.On("DeleteAddressByUserId", 3, 1).Return(fmt.Errorf("error"))
	err := usecase.DeleteAddressByUserId(3, 1)

	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestAddressUsecaseImpl_DeleteAddressByUserId_Success(t *testing.T) {
	mockAddressRepo := mocks.NewAddressRepository(t)
	usecase := NewAddressUsecaseImpl(mockAddressRepo)

	mockAddressRepo.On("DeleteAddressByUserId", 3, 1).Return(nil)
	err := usecase.DeleteAddressByUserId(3, 1)

	assert.Equal(t, nil, err)
}
