package usecase

import (
	"final-project-backend/dto"
)

type AddressUsecase interface {
	GetAllAddress(page int, limit int, search string) (dto.AddressPaginateResponse, error)
	GetAllAddressByUserId(userId int, page int, limit int, search string) (dto.AddressPaginateResponse, error)
	GetAddressByUserId(userId int, addressId int) (*dto.AddressResponse, error)
	CreateAddress(request dto.AddressCreateRequest) (*dto.AddressResponse, error)
	UpdateAddressByUserId(request dto.AddressUpdateRequest) (*dto.AddressResponse, error)
	DeleteAddressByUserId(userId int, addressId int) error
}
