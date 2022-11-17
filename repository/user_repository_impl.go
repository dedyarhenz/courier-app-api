package repository

import (
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) GetUserById(userId int) (*entity.User, error) {
	return nil, nil
}

func (r *UserRepositoryImpl) GetUserByEmail(email string) (*entity.User, error) {
	var usr entity.User

	if err := r.db.Where("email ILIKE ?", email).First(&usr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custErr.ErrUserNotFound
		}

		return nil, err
	}

	return &usr, nil
}

func (r *UserRepositoryImpl) CreateUser(user entity.User) (*entity.User, error) {
	usr := entity.User{
		Email:        user.Email,
		Password:     user.Password,
		FullName:     user.FullName,
		Phone:        user.Phone,
		Role:         user.Role,
		Balance:      user.Balance,
		Photo:        user.Photo,
		RefferalCode: user.RefferalCode,
	}

	if err := r.db.Omit("created_at", "updated_at", "deleted_at").Create(&usr).Error; err != nil {
		return nil, err
	}

	return &usr, nil
}
