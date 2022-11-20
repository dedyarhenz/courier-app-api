package repository

import "final-project-backend/entity"

type SizeRepository interface {
	GetAllSize() ([]entity.Size, error)
	GetSizeById(sizeId int) (*entity.Size, error)
}
