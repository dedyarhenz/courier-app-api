package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"
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
	if userAlready != nil {
		return nil, custErr.ErrEmailAlready
	}

	user := entity.User{
		Email:    request.Email,
		FullName: request.FullName,
		Phone:    request.Phone,
		Photo:    request.Phone,
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

	user, err := u.repoUser.AddBalance(request.UserId, request.Amount)
	if err != nil {
		return nil, err
	}

	resUser := dto.CreateUserResponse(*user)

	return &resUser, nil
}
