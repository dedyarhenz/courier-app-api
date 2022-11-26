package dto

type PaymentPayRequest struct {
	PaymentId   int  `json:"-"`
	UserId      int  `json:"-"`
	PromoUserId *int `json:"promo_user_id" binding:"omitempty,numeric"`
}
