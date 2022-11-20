package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/repository"
)

type SizeUsecaseImpl struct {
	repoSize repository.SizeRepository
}

func NewSizeUsecaseImpl(repoSize repository.SizeRepository) SizeUsecase {
	return &SizeUsecaseImpl{
		repoSize: repoSize,
	}
}

func (u *SizeUsecaseImpl) GetAllSize() ([]dto.SizeResponse, error) {
	allSize, err := u.repoSize.GetAllSize()
	if err != nil {
		return nil, err
	}

	resAllSize := dto.CreateSizeListResponse(allSize)

	return resAllSize, nil
}
