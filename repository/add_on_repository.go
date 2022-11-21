package repository

import "final-project-backend/entity"

type AddOnRepository interface {
	GetAllAddOn() ([]entity.AddOn, error)
	GetAddOnById(addOnId int) (*entity.AddOn, error)
	GetAddOnByMultipleId(addOnsId []int) ([]entity.AddOn, error)
}
