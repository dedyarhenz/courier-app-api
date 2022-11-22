package dto

type ShippingReviewRequest struct {
	UserId     int    `json:"-"`
	ShippingId int    `json:"-"`
	Review     string `json:"review" binding:"required"`
}
