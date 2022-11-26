package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/repository"
)

type PromoUserUsecaseImpl struct {
	repoPromoUser repository.PromoUserRepository
}

func NewPromoUserUsecaseImpl(repoPromoUser repository.PromoUserRepository) PromoUserUsecase {
	return &PromoUserUsecaseImpl{
		repoPromoUser: repoPromoUser,
	}
}

func (u *PromoUserUsecaseImpl) GetAllPromoUserByUserId(userId int) ([]dto.PromoUserResponse, error) {
	promoUsers, err := u.repoPromoUser.GetAllPromoUserByUserId(userId)
	if err != nil {
		return nil, err
	}

	resPromoUser := dto.CreatePromoUserListResponse(promoUsers)

	return resPromoUser, nil
}
