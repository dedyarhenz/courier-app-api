package dto

import "final-project-backend/entity"

type PaymentResponse struct {
	Id            int    `json:"id"`
	PaymentStatus string `json:"payment_status"`
	TotalCost     int    `json:"total_cost"`
	PromoId       *int   `json:"promo_id"`
}

func CreatePaymentResponse(payment entity.Payment) PaymentResponse {
	return PaymentResponse{
		Id:            payment.Id,
		PaymentStatus: payment.PaymentStatus,
		TotalCost:     payment.TotalCost,
		PromoId:       payment.PromoId,
	}
}
