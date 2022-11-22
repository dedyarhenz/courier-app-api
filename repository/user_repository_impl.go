package repository

import (
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	var usr entity.User

	if err := r.db.First(&usr, userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custErr.ErrUserNotFound
		}

		return nil, err
	}

	return &usr, nil
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

func (r *UserRepositoryImpl) GetUserByRefferalCode(refferalCode string) (*entity.User, error) {
	var usr entity.User

	if err := r.db.Where("refferal_code = ?", refferalCode).First(&usr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custErr.ErrReffCodeInvalid
		}

		return nil, err
	}

	return &usr, nil
}

func (r *UserRepositoryImpl) CreateUser(user entity.User) (*entity.User, error) {
	usr := entity.User{
		Email:          user.Email,
		Password:       user.Password,
		FullName:       user.FullName,
		Phone:          user.Phone,
		Role:           user.Role,
		Balance:        user.Balance,
		Photo:          user.Photo,
		RefferalCode:   user.RefferalCode,
		RefferedUserId: user.RefferedUserId,
	}

	if err := r.db.Omit("created_at", "updated_at", "deleted_at").Create(&usr).Error; err != nil {
		return nil, err
	}

	return &usr, nil
}

func (r *UserRepositoryImpl) UpdateUser(user entity.User) (*entity.User, error) {
	usr := entity.User{
		Email:    user.Email,
		FullName: user.FullName,
		Phone:    user.Phone,
		Photo:    user.Photo,
	}

	res := r.db.
		Omit("created_at", "updated_at", "deleted_at").
		Clauses(clause.Returning{}).
		Where("id = ?", user.Id).
		Updates(&usr)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, custErr.ErrUserNotFound
	}

	return &usr, nil
}

func (r *UserRepositoryImpl) AddBalance(userId int, amount int) (*entity.User, error) {
	var usr entity.User

	res := r.db.
		Clauses(clause.Returning{}).
		Model(&usr).Where("id = ?", userId).
		UpdateColumn("balance", gorm.Expr("balance + ?", amount))

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, custErr.ErrUserNotFound
	}

	return &usr, nil
}

func (r *UserRepositoryImpl) ReduceBalance(userId int, amount int) (*entity.User, error) {
	var usr entity.User

	res := r.db.
		Clauses(clause.Returning{}).
		Model(&usr).Where("id = ?", userId).
		UpdateColumn("balance", gorm.Expr("balance - ?", amount))

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, custErr.ErrUserNotFound
	}

	return &usr, nil
}
