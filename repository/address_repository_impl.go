package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type AddressRepositoryImpl struct {
	db *gorm.DB
}

func NewAddressRepositoryImpl(db *gorm.DB) AddressRepository {
	return &AddressRepositoryImpl{
		db: db,
	}
}

func (r *AddressRepositoryImpl) GetAllAddress() ([]entity.Address, error) {
	var addresses []entity.Address

	err := r.db.Find(&addresses).Error

	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (r *AddressRepositoryImpl) GetAddressByUserId(userId int) ([]entity.Address, error) {
	var addresses []entity.Address

	err := r.db.Where("user_id = ?", userId).Find(&addresses).Error

	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (r *AddressRepositoryImpl) CreateAddress(address entity.Address) (*entity.Address, error) {
	newAddress := entity.Address{
		RecipientName:  address.RecipientName,
		FullAddress:    address.FullAddress,
		RecipientPhone: address.RecipientPhone,
		UserId:         address.UserId,
	}

	if err := r.db.Omit("created_at", "updated_at", "deleted_at").Create(&newAddress).Error; err != nil {
		return nil, err
	}

	return &newAddress, nil
}
