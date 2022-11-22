package dto

type PaymentPayRequest struct {
	PaymentId int  `json:"-"`
	UserId    int  `json:"-"`
	PromoId   *int `json:"promo_id" binding:"omitempty,numeric"`
}
