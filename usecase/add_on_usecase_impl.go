package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/repository"
)

type AddOnUsecaseImpl struct {
	repoAddOn repository.AddOnRepository
}

func NewAddOnUsecaseImpl(repoAddOn repository.AddOnRepository) AddOnUsecase {
	return &AddOnUsecaseImpl{
		repoAddOn: repoAddOn,
	}
}

func (u *AddOnUsecaseImpl) GetAllAddOn() ([]dto.AddOnResponse, error) {
	allAddOn, err := u.repoAddOn.GetAllAddOn()
	if err != nil {
		return nil, err
	}

	resAllAddOn := dto.CreateAddOnListResponse(allAddOn)

	return resAllAddOn, nil
}
