package repository

import (
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PromoUserRepositoryImpl struct {
	db *gorm.DB
}

func NewPromoUserRepositoryImpl(db *gorm.DB) PromoUserRepository {
	return &PromoUserRepositoryImpl{
		db: db,
	}
}

func (r *PromoUserRepositoryImpl) GetPromoUserById(promoUserId int) (*entity.PromoUser, error) {
	var promoUser entity.PromoUser
	err := r.db.Unscoped().Preload("Promo").First(&promoUser, promoUserId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custErr.ErrPromoUserNotFound
		}

		return nil, err
	}

	return &promoUser, nil
}

func (r *PromoUserRepositoryImpl) UpdatePromoUser(promo entity.PromoUser) (*entity.PromoUser, error) {
	newPromoUser := entity.PromoUser{}
	res := r.db.
		Clauses(clause.Returning{}).
		Model(&newPromoUser).
		Omit("created_at", "updated_at", "deleted_at").
		Where("id = ?", promo.Id).
		Updates(promo)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, custErr.ErrPromoUserNotFound
	}

	return &newPromoUser, nil
}

func (r *PromoUserRepositoryImpl) CreatePromoUser(promo entity.PromoUser) (*entity.PromoUser, error) {
	newPromoUser := entity.PromoUser{
		PromoId: promo.PromoId,
		UserId:  promo.UserId,
		IsUsed:  promo.IsUsed,
	}

	err := r.db.Omit("created_at", "updated_at", "deleted_at").Create(&newPromoUser).Error
	if err != nil {
		return nil, err
	}

	return &newPromoUser, nil
}

func (r *PromoUserRepositoryImpl) GetAllPromoUserByUserId(userId int) ([]entity.PromoUser, error) {
	var promoUsers []entity.PromoUser
	err := r.db.Unscoped().Preload("Promo").Where("user_id", userId).Find(&promoUsers).Error
	if err != nil {

		return nil, err
	}

	return promoUsers, nil
}
