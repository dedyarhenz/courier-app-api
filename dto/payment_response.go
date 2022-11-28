package dto

import "final-project-backend/entity"

type PaymentResponse struct {
	Id            int            `json:"id"`
	PaymentStatus string         `json:"payment_status"`
	TotalCost     int            `json:"total_cost"`
	PromoId       *int           `json:"promo_id"`
	Promo         *PromoResponse `json:"promo"`
}

func CreatePaymentResponse(payment entity.Payment) PaymentResponse {
	var promo *PromoResponse
	if payment.Promo != nil {
		res := CreatePromoResponse(*payment.Promo)
		promo = &res
	}

	return PaymentResponse{
		Id:            payment.Id,
		PaymentStatus: payment.PaymentStatus,
		TotalCost:     payment.TotalCost,
		PromoId:       payment.PromoId,
		Promo:         promo,
	}
}
