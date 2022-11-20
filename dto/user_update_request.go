package dto

type UserUpdateRequest struct {
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Photo    string `json:"photo" binding:"omitempty"`
}
