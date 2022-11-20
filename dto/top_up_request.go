package dto

type TopUpRequest struct {
	Amount int `json:"amount" binding:"required,numeric"`
	UserId int `json:"-"`
}
