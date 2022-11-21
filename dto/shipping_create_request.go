package dto

type ShippingCreateRequest struct {
	UserId     int   `json:"-"`
	SizeId     int   `json:"size_id" binding:"required,numeric"`
	CategoryId int   `json:"category_id" binding:"required,numeric"`
	AddressId  int   `json:"address_id" binding:"required,numeric"`
	AddOnsId   []int `json:"add_ons_id" binding:"required,gt=0"`
}
