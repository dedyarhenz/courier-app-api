package usecase

import (
	"final-project-backend/dto"
)

type AuthUsecase interface {
	Login(request dto.UserLoginRequest) (string, error)
	Register(request dto.UserRegisterRequest) (*dto.UserRegisterResponse, error)
}
