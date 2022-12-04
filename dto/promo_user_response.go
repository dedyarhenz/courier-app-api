package dto

import "final-project-backend/entity"

type PromoUserResponse struct {
	Id     int            `json:"id"`
	IsUsed bool           `json:"is_used"`
	Promo  *PromoResponse `json:"promo"`
}

func CreatePromoUserResponse(promoUser entity.PromoUser) PromoUserResponse {
	var promo *PromoResponse
	if promoUser.Promo != nil {
		res := CreatePromoResponse(*promoUser.Promo)
		promo = &res
	}

	return PromoUserResponse{
		Id:    promoUser.Id,
		Promo: promo,
	}
}

func CreatePromoUserListResponse(promoUsers []entity.PromoUser) []PromoUserResponse {
	promoUsersResponse := []PromoUserResponse{}
	for _, p := range promoUsers {
		promoUser := CreatePromoUserResponse(p)
		promoUsersResponse = append(promoUsersResponse, promoUser)
	}

	return promoUsersResponse
}
