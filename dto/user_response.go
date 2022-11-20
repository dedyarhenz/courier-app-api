package dto

import "final-project-backend/entity"

type UserResponse struct {
	Id           int    `gorm:"primaryKey;column:id"`
	Email        string `gorm:"column:email"`
	FullName     string `gorm:"column:long_name"`
	Phone        string `gorm:"column:phone"`
	Balance      int    `gorm:"column:balance"`
	Photo        string `gorm:"column:photo"`
	RefferalCode string `gorm:"column:refferal_code"`
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
