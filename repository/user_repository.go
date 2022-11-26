package repository

import "final-project-backend/entity"

type UserRepository interface {
	GetUserById(userId int) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUserByRefferalCode(refferalCode string) (*entity.User, error)
	CreateUser(user entity.User) (*entity.User, error)
	UpdateUser(user entity.User) (*entity.User, error)
	AddBalance(userId int, amount int) (*entity.User, error)
	ReduceBalance(userId int, amount int) (*entity.User, error)
	UpdatedCompleteBonus(userId int, completeBonus bool) error
	UpdatedCompleteBonusReff(userId int, completeBonusReff bool) error
}
