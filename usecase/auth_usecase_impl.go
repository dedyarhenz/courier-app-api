package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"
	"final-project-backend/pkg/utils"
	"final-project-backend/repository"
)

type AuthUsecaseImp struct {
	repoUser repository.UserRepository
}

func NewAuthUsecaseImpl(repoUser repository.UserRepository) AuthUsecase {
	return &AuthUsecaseImp{
		repoUser: repoUser,
	}
}

func (u *AuthUsecaseImp) Login(request dto.UserLoginRequest) (string, error) {
	user, err := u.repoUser.GetUserByEmail(request.Email)
	if err != nil {
		return "", err
	}

	ok := utils.ComparePassword(user.Password, request.Password)
	if !ok {
		return "", custErr.ErrLoginFailed
	}

	token, err := utils.GenerateJWT(user.Id, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *AuthUsecaseImp) Register(request dto.UserRegisterRequest) error {
	userAlready, _ := u.repoUser.GetUserByEmail(request.Email)
	if userAlready != nil {
		return custErr.ErrEmailAlready
	}

	hashPass, err := utils.HashAndSalt(request.Password)
	if err != nil {
		return err
	}

	userNew := entity.User{
		Email:        request.Email,
		Password:     hashPass,
		FullName:     request.FullName,
		Phone:        request.Phone,
		Role:         entity.UserRole,
		Balance:      0,
		RefferalCode: "",
	}

	_, err = u.repoUser.CreateUser(userNew)
	if err != nil {
		return err
	}

	return nil
}
