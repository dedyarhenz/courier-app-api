package usecase

import "final-project-backend/dto"

type PromoUsecase interface {
	GetAllPromo(page int, limit int, search string, order string, sort string) (dto.PromoPaginateResponse, error)
	GetPromoById(promoId int) (*dto.PromoResponse, error)
	CreatePromo(request dto.PromoCreateRequest) (*dto.PromoResponse, error)
	UpdatePromo(request dto.PromoUpdateRequest) (*dto.PromoResponse, error)
}
