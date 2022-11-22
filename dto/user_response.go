package dto

import "final-project-backend/entity"

type UserResponse struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	Phone        string `json:"phone"`
	Balance      int    `json:"balance"`
	Photo        string `json:"photo"`
	RefferalCode string `json:"refferal_code"`
}

func CreateUserResponse(user entity.User) UserResponse {
	return UserResponse{
		Id:           user.Id,
		Email:        user.Email,
		FullName:     user.FullName,
		Phone:        user.Phone,
		Balance:      user.Balance,
		Photo:        user.Photo,
		RefferalCode: user.RefferalCode,
	}
}

func CreateUserListResponse(users []entity.User) []UserResponse {
	usersResponse := []UserResponse{}
	for _, u := range users {
		user := CreateUserResponse(u)
		usersResponse = append(usersResponse, user)
	}

	return usersResponse
}
