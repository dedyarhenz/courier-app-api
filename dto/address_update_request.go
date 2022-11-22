package dto

type AddressUpdateRequest struct {
	Id             int    `json:"-"`
	RecipientName  string `json:"recipient_name" binding:"required"`
	FullAddress    string `json:"full_address" binding:"required"`
	RecipientPhone string `json:"recipient_phone" binding:"required"`
	UserId         int    `json:"-"`
}
