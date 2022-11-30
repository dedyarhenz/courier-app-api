package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"
	"final-project-backend/pkg/helper"
	"final-project-backend/repository"
)

type UserUsecaseImpl struct {
	repoUser repository.UserRepository
}

func NewUserUsecaseImpl(repoUser repository.UserRepository) UserUsecase {
	return &UserUsecaseImpl{
		repoUser: repoUser,
	}
}

func (u *UserUsecaseImpl) GetUserById(userId int) (*dto.UserResponse, error) {
	user, err := u.repoUser.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	resUser := dto.CreateUserResponse(*user)

	return &resUser, nil
}

func (u *UserUsecaseImpl) UpdateUserById(request dto.UserUpdateRequest) (*dto.UserResponse, error) {
	userAlready, _ := u.repoUser.GetUserByEmail(request.Email)
	if userAlready != nil && userAlready.Email != request.Email {
		return nil, custErr.ErrEmailAlready
	}

	var photoUrl string
	var err error
	if request.Photo != nil {
		photoUrl, err = helper.ImageUploadHelper(request.Photo)
		if err != nil {
			return nil, err
		}
	}

	user := entity.User{
		Id:       request.Id,
		Email:    request.Email,
		FullName: request.FullName,
		Phone:    request.Phone,
		Photo:    photoUrl,
	}

	userUpdated, err := u.repoUser.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	resUserUpdated := dto.CreateUserResponse(*userUpdated)

	return &resUserUpdated, nil
}

func (u *UserUsecaseImpl) TopUp(request dto.TopUpRequest) (*dto.UserResponse, error) {
	if request.Amount < entity.MinTopUp {
		return nil, custErr.ErrMinTopUp
	}

	if request.Amount > entity.MaxTopUp {
		return nil, custErr.ErrMaxTopUp
	}

	user, err := u.repoUser.AddBalance(request.UserId, request.Amount)
	if err != nil {
		return nil, err
	}

	resUser := dto.CreateUserResponse(*user)

	return &resUser, nil
}
