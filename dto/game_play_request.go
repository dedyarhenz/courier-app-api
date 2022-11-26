package dto

type GamePlayRequest struct {
	UserId     int `json:"-"`
	ShippingId int `json:"shipping_id" binding:"required"`
}
