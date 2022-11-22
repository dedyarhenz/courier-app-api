package repository

import (
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *AddressRepositoryImpl) GetAddressBySpecificUser(userId int, addresId int) (*entity.Address, error) {
	var address entity.Address

	err := r.db.Where("user_id = ?", userId).First(&address, addresId).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custErr.ErrAddressNotFound
		}
		return nil, err
	}

	return &address, nil
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

func (r *AddressRepositoryImpl) UpdateAddressByUserId(address entity.Address) (*entity.Address, error) {
	newAddress := entity.Address{
		RecipientName:  address.RecipientName,
		FullAddress:    address.FullAddress,
		RecipientPhone: address.RecipientPhone,
		UserId:         address.UserId,
	}
	res := r.db.
		Clauses(clause.Returning{}).
		Omit("created_at", "updated_at", "deleted_at").
		Where("user_id = ?", address.UserId).
		Where("id = ?", address.Id).
		UpdateColumns(&newAddress)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, custErr.ErrAddressNotFound
	}

	return &newAddress, nil
}
