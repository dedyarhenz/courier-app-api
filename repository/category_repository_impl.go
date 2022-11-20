package repository

import (
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

func (r *CategoryRepositoryImpl) GetAllCategory() ([]entity.Category, error) {
	var categories []entity.Category

	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepositoryImpl) GetCategoryById(categoryId int) (*entity.Category, error) {
	var category entity.Category

	err := r.db.First(&category, categoryId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custErr.ErrCategoryNotFound
		}

		return nil, err
	}

	return &category, nil
}
