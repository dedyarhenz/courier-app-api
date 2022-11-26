package usecase

import "final-project-backend/dto"

type PromoUserUsecase interface {
	GetAllPromoUserByUserId(userId int) ([]dto.PromoUserResponse, error)
}
