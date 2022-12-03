package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"
	"final-project-backend/pkg/utils"
	"final-project-backend/repository"
	"math/rand"
	"time"
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

	token, err := utils.GenerateJWT(user.Id, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *AuthUsecaseImp) Register(request dto.UserRegisterRequest) (*dto.UserRegisterResponse, error) {
	userAlready, _ := u.repoUser.GetUserByEmail(request.Email)
	if userAlready != nil {
		return nil, custErr.ErrEmailAlready
	}

	hashPass, err := utils.HashAndSalt(request.Password)
	if err != nil {
		return nil, err
	}

	var reffId *int
	if request.RefferalCode != "" {
		userReff, err := u.repoUser.GetUserByRefferalCode(request.RefferalCode)
		if err != nil {
			return nil, err
		}
		reffId = &userReff.Id
	}

	userNew := entity.User{
		Email:          request.Email,
		Password:       hashPass,
		FullName:       request.FullName,
		Phone:          request.Phone,
		Role:           entity.UserRole,
		Balance:        0,
		RefferalCode:   randomString(8),
		RefferedUserId: reffId,
	}

	user, err := u.repoUser.CreateUser(userNew)
	if err != nil {
		return nil, err
	}

	userRes := dto.UserRegisterResponse{
		Email:        user.Email,
		FullName:     user.FullName,
		Phone:        user.Phone,
		RefferalCode: user.RefferalCode,
	}

	return &userRes, nil
}

func randomString(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
