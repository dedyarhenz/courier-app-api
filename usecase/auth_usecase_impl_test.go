package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	custErr "final-project-backend/pkg/errors"
	"final-project-backend/pkg/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthUsecaseImp_Login_ErrorEmailAlreadyExist(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewAuthUsecaseImpl(mockUserRepo)

	userLoginRequest := dto.UserLoginRequest{
		Email:    "dedy@gmail.com",
		Password: "123456",
	}

	mockUserRepo.On("GetUserByEmail", userLoginRequest.Email).Return(nil, custErr.ErrEmailAlready)
	res, err := usecase.Login(userLoginRequest)

	assert.NotNil(t, res)
	assert.Equal(t, custErr.ErrEmailAlready, err)
}

func TestAuthUsecaseImp_Login_ErrorPassword(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewAuthUsecaseImpl(mockUserRepo)

	userLoginRequest := dto.UserLoginRequest{
		Email:    "dedy@gmail.com",
		Password: "123456",
	}

	user := entity.User{
		Id:       1,
		Email:    "dedy@gmail.com",
		Password: "asdasd",
		Role:     "USER",
	}

	mockUserRepo.On("GetUserByEmail", userLoginRequest.Email).Return(&user, nil)
	utils.ComparePassword(user.Password, userLoginRequest.Password)

	res, err := usecase.Login(userLoginRequest)

	assert.NotNil(t, res)
	assert.Equal(t, custErr.ErrLoginFailed, err)
}

func TestAuthUsecaseImp_Login_Success(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewAuthUsecaseImpl(mockUserRepo)

	userLoginRequest := dto.UserLoginRequest{
		Email:    "dedy@gmail.com",
		Password: "123456",
	}

	passHash, _ := utils.HashAndSalt(userLoginRequest.Password)

	user := entity.User{
		Id:       1,
		Email:    "dedy@gmail.com",
		Password: passHash,
		Role:     "USER",
	}

	mockUserRepo.On("GetUserByEmail", userLoginRequest.Email).Return(&user, nil)
	utils.ComparePassword(user.Password, userLoginRequest.Password)
	token, _ := utils.GenerateJWT(user.Id, user.Role)

	res, err := usecase.Login(userLoginRequest)

	assert.Nil(t, err)
	assert.Equal(t, token, res)
}

func TestAuthUsecaseImp_Register_ErrorEmailAlreadyExist(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewAuthUsecaseImpl(mockUserRepo)

	userRegisterRequest := dto.UserRegisterRequest{
		Email:        "dedy@gmail.com",
		Password:     "123456",
		FullName:     "dedy",
		Phone:        "0987654",
		RefferalCode: "qwerty12",
	}
	user := entity.User{
		Email:        "dedy@gmail.com",
		Password:     "123456",
		FullName:     "dedy",
		Phone:        "0987654",
		RefferalCode: "qwerty12",
	}

	mockUserRepo.On("GetUserByEmail", userRegisterRequest.Email).Return(&user, nil)
	res, err := usecase.Register(userRegisterRequest)

	assert.Nil(t, res)
	assert.Equal(t, custErr.ErrEmailAlready, err)
}

func TestAuthUsecaseImp_Register_ErrorRefferal(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewAuthUsecaseImpl(mockUserRepo)

	userRegisterRequest := dto.UserRegisterRequest{
		Email:        "dedy@gmail.com",
		Password:     "123456",
		FullName:     "dedy",
		Phone:        "0987654",
		RefferalCode: "qwerty12",
	}

	mockUserRepo.On("GetUserByEmail", userRegisterRequest.Email).Return(nil, fmt.Errorf("error"))
	mockUserRepo.On("GetUserByRefferalCode", userRegisterRequest.RefferalCode).Return(nil, fmt.Errorf("error"))

	res, err := usecase.Register(userRegisterRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}
