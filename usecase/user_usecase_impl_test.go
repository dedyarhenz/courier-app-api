package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	custErr "final-project-backend/pkg/errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserUsecaseImpl_GetUserById_Error(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewUserUsecaseImpl(mockUserRepo)

	mockUserRepo.On("GetUserById", 1).Return(nil, fmt.Errorf("error"))
	res, err := usecase.GetUserById(1)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestUserUsecaseImpl_GetUserById_Success(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewUserUsecaseImpl(mockUserRepo)

	user := entity.User{
		Id:       1,
		Email:    "dedy@gmail.com",
		Password: "123456",
		FullName: "dedy",
		Phone:    "098776",
		Role:     "USER",
		Balance:  0,
	}

	userResponse := dto.CreateUserResponse(user)

	mockUserRepo.On("GetUserById", 1).Return(&user, nil)
	res, err := usecase.GetUserById(1)

	assert.Nil(t, err)
	assert.Equal(t, &userResponse, res)
}

func TestUserUsecaseImpl_UpdateUserById_ErrorUserNotFound(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewUserUsecaseImpl(mockUserRepo)

	userRequest := dto.UserUpdateRequest{
		Id:       1,
		Email:    "ded@gmail.com",
		FullName: "dedy",
		Phone:    "098776",
		Photo:    nil,
	}

	mockUserRepo.On("GetUserById", userRequest.Id).Return(nil, fmt.Errorf("error"))
	res, err := usecase.UpdateUserById(userRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestUserUsecaseImpl_UpdateUserById_ErrorEmailAlreadyExist(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewUserUsecaseImpl(mockUserRepo)

	userRequest := dto.UserUpdateRequest{
		Id:       1,
		Email:    "ded@gmail.com",
		FullName: "dedy",
		Phone:    "098776",
		Photo:    nil,
	}

	userOld := entity.User{
		Id:       1,
		Email:    "dedy@gmail.com",
		Password: "123456",
		FullName: "dedy",
		Phone:    "098776",
		Role:     "USER",
		Balance:  0,
	}

	user := entity.User{
		Id:       1,
		Email:    "ded@gmail.com",
		Password: "123456",
		FullName: "dedy",
		Phone:    "098776",
		Role:     "USER",
		Balance:  0,
	}

	mockUserRepo.On("GetUserById", userRequest.Id).Return(&userOld, nil)
	mockUserRepo.On("GetUserByEmail", userRequest.Email).Return(&user, nil)
	res, err := usecase.UpdateUserById(userRequest)

	assert.Nil(t, res)
	assert.Equal(t, custErr.ErrEmailAlready, err)
}

func TestUserUsecaseImpl_UpdateUserById_ErrorUpdate(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewUserUsecaseImpl(mockUserRepo)

	userRequest := dto.UserUpdateRequest{
		Id:       1,
		Email:    "ded@gmail.com",
		FullName: "dedy",
		Phone:    "098776",
		Photo:    nil,
	}

	userOld := entity.User{
		Id:       1,
		Email:    "dedy@gmail.com",
		Password: "123456",
		FullName: "dedy",
		Phone:    "098776",
		Role:     "USER",
		Balance:  0,
	}

	user := entity.User{
		Id:       1,
		Email:    "ded@gmail.com",
		FullName: "dedy",
		Phone:    "098776",
		Balance:  0,
	}

	mockUserRepo.On("GetUserById", userRequest.Id).Return(&userOld, nil)
	mockUserRepo.On("GetUserByEmail", userRequest.Email).Return(nil, fmt.Errorf("error"))
	mockUserRepo.On("UpdateUser", user).Return(nil, fmt.Errorf("error"))
	res, err := usecase.UpdateUserById(userRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestUserUsecaseImpl_UpdateUserById_Success(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewUserUsecaseImpl(mockUserRepo)

	userRequest := dto.UserUpdateRequest{
		Id:       1,
		Email:    "ded@gmail.com",
		FullName: "dedy",
		Phone:    "098776",
		Photo:    nil,
	}

	userOld := entity.User{
		Id:       1,
		Email:    "dedy@gmail.com",
		Password: "123456",
		FullName: "dedy",
		Phone:    "098776",
		Role:     "USER",
		Balance:  0,
	}

	user := entity.User{
		Id:       1,
		Email:    "ded@gmail.com",
		FullName: "dedy",
		Phone:    "098776",
		Balance:  0,
	}

	userResponse := dto.CreateUserResponse(user)

	mockUserRepo.On("GetUserById", userRequest.Id).Return(&userOld, nil)
	mockUserRepo.On("GetUserByEmail", userRequest.Email).Return(nil, fmt.Errorf("error"))
	mockUserRepo.On("UpdateUser", user).Return(&user, nil)
	res, err := usecase.UpdateUserById(userRequest)

	assert.Nil(t, err)
	assert.Equal(t, &userResponse, res)
}

func TestUserUsecaseImpl_TopUp_ErrorMinAmount(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewUserUsecaseImpl(mockUserRepo)

	topUpRequest := dto.TopUpRequest{
		Amount: 1000,
		UserId: 1,
	}

	res, err := usecase.TopUp(topUpRequest)

	assert.Nil(t, res)
	assert.Equal(t, custErr.ErrMinTopUp, err)
}

func TestUserUsecaseImpl_TopUp_ErrorMaxAmount(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewUserUsecaseImpl(mockUserRepo)

	topUpRequest := dto.TopUpRequest{
		Amount: 20000000,
		UserId: 1,
	}

	res, err := usecase.TopUp(topUpRequest)

	assert.Nil(t, res)
	assert.Equal(t, custErr.ErrMaxTopUp, err)
}

func TestUserUsecaseImpl_TopUp_ErrorAddBalance(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewUserUsecaseImpl(mockUserRepo)

	topUpRequest := dto.TopUpRequest{
		Amount: 300000,
		UserId: 1,
	}

	mockUserRepo.On("AddBalance", topUpRequest.UserId, topUpRequest.Amount).Return(nil, fmt.Errorf("error"))

	res, err := usecase.TopUp(topUpRequest)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestUserUsecaseImpl_TopUp_Success(t *testing.T) {
	mockUserRepo := mocks.NewUserRepository(t)
	usecase := NewUserUsecaseImpl(mockUserRepo)

	topUpRequest := dto.TopUpRequest{
		Amount: 300000,
		UserId: 1,
	}

	user := entity.User{
		Id:      1,
		Email:   "dedy@gmail.com",
		Balance: 500000,
	}

	userResponse := dto.CreateUserResponse(user)

	mockUserRepo.On("AddBalance", topUpRequest.UserId, topUpRequest.Amount).Return(&user, nil)

	res, err := usecase.TopUp(topUpRequest)

	assert.Nil(t, err)
	assert.Equal(t, &userResponse, res)
}
