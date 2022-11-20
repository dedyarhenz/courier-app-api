package repository

import (
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"

	"gorm.io/gorm"
)

type SizeRepositoryImpl struct {
	db *gorm.DB
}

func NewSizeRepositoryImpl(db *gorm.DB) SizeRepository {
	return &SizeRepositoryImpl{
		db: db,
	}
}

func (r *SizeRepositoryImpl) GetAllSize() ([]entity.Size, error) {
	var sizes []entity.Size

	err := r.db.Find(&sizes).Error
	if err != nil {
		return nil, err
	}

	return sizes, nil
}

func (r *SizeRepositoryImpl) GetSizeById(sizeId int) (*entity.Size, error) {
	var size entity.Size

	err := r.db.First(&size, sizeId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custErr.ErrSizeNotFound
		}

		return nil, err
	}

	return nil, nil
}
