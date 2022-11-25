package dto

import (
	"final-project-backend/entity"
	"time"
)

type PromoResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	MinFee      int       `json:"min_fee"`
	Discount    int       `json:"discount"`
	MaxDiscount int       `json:"max_discount"`
	Quota       int       `json:"quota"`
	ExpireDate  time.Time `json:"expire_date"`
}

func CreatePromoResponse(promo entity.Promo) PromoResponse {
	return PromoResponse{
		Id:          promo.Id,
		Name:        promo.Name,
		MinFee:      promo.MinFee,
		Discount:    promo.Discount,
		MaxDiscount: promo.MaxDiscount,
		Quota:       promo.Quota,
		ExpireDate:  promo.ExpireDate,
	}
}

func CreatePromoListResponse(promos []entity.Promo) []PromoResponse {
	promoResponse := []PromoResponse{}
	for _, c := range promos {
		promo := CreatePromoResponse(c)
		promoResponse = append(promoResponse, promo)
	}

	return promoResponse
}
