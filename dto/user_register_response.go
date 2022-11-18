package dto

type UserRegisterResponse struct {
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	Phone        string `json:"phone"`
	RefferalCode string `json:"refferal_code"`
}
