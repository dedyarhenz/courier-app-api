package repository

import "final-project-backend/entity"

type CategoryRepository interface {
	GetAllCategory() ([]entity.Category, error)
	GetCategoryById(categoryId int) (*entity.Category, error)
}
