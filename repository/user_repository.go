package repository

import "final-project-backend/entity"

type UserRepository interface {
	GetUserById(userId int) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	CreateUser(user entity.User) (*entity.User, error)
}
