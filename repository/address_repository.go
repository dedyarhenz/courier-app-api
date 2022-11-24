package repository

import "final-project-backend/entity"

type AddressRepository interface {
	GetAllAddress() ([]entity.Address, error)
	GetAllAddressByUserId(userId int) ([]entity.Address, error)
	GetAddressByUserId(userId int, addressId int) (*entity.Address, error)
	CreateAddress(address entity.Address) (*entity.Address, error)
	UpdateAddressByUserId(address entity.Address) (*entity.Address, error)
}
