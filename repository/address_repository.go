package repository

import "final-project-backend/entity"

type AddressRepository interface {
	GetAllAddress(offset int, limit int, search string, orderAndSort string) ([]entity.Address, error)
	GetAllAddressByUserId(userId int, offset int, limit int, search string, orderAndSort string) ([]entity.Address, error)
	GetAddressByUserId(userId int, addressId int) (*entity.Address, error)
	CreateAddress(address entity.Address) (*entity.Address, error)
	UpdateAddressByUserId(address entity.Address) (*entity.Address, error)
	CountAddress() int64
	CountAddressByUserId(userId int) int64
}
