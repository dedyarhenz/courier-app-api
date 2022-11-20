package usecase

import "final-project-backend/dto"

type UserUsecase interface {
	GetUserById(userId int) (*dto.UserResponse, error)
	UpdateUserById(request dto.UserUpdateRequest) (*dto.UserResponse, error)
	TopUp(request dto.TopUpRequest) (*dto.UserResponse, error)
}
