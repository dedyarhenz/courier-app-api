package dto

type UserRegisterRequest struct {
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
	FullName     string `json:"full_name" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	RefferalCode string `json:"refferal_code" binding:"omitempty,len=8"`
}
