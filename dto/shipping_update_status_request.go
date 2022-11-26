package dto

type ShippingUpdateStatusRequest struct {
	ShippingId     int    `json:"-"`
	StatusShipping string `json:"status_shipping" binding:"required"`
}
