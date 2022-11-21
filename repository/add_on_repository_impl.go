package repository

import (
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"

	"gorm.io/gorm"
)

type AddOnRepositoryImpl struct {
	db *gorm.DB
}

func NewAddOnRepositoryImpl(db *gorm.DB) AddOnRepository {
	return &AddOnRepositoryImpl{
		db: db,
	}
}

func (r *AddOnRepositoryImpl) GetAllAddOn() ([]entity.AddOn, error) {
	var addOns []entity.AddOn

	err := r.db.Find(&addOns).Error
	if err != nil {
		return nil, err
	}

	return addOns, nil
}

func (r *AddOnRepositoryImpl) GetAddOnById(addOnId int) (*entity.AddOn, error) {
	var addOn entity.AddOn

	err := r.db.First(&addOn, addOnId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custErr.ErrAddOnNotFound
		}

		return nil, err
	}

	return &addOn, nil
}

func (r *AddOnRepositoryImpl) GetAddOnByMultipleId(addOnsId []int) ([]entity.AddOn, error) {
	var addOns []entity.AddOn

	err := r.db.Find(&addOns, addOnsId).Error
	if err != nil {
		return nil, err
	}

	return addOns, nil
}
