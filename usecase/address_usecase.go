package usecase

import (
	"final-project-backend/dto"
)

type AddressUsecase interface {
	GetAllAddress() ([]dto.AddressResponse, error)
	GetAddressByUserId(userId int) ([]dto.AddressResponse, error)
	CreateAddress(request dto.AddressCreateRequest) (*dto.AddressResponse, error)
	UpdateAddressByUserId(request dto.AddressUpdateRequest) (*dto.AddressResponse, error)
}
