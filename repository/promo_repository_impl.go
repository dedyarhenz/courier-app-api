package repository

import (
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PromoRepositoryImpl struct {
	db *gorm.DB
}

func NewPromoRepositoryImpl(db *gorm.DB) PromoRepository {
	return &PromoRepositoryImpl{
		db: db,
	}
}

func (r *PromoRepositoryImpl) GetPromoById(promoId int) (*entity.Promo, error) {
	var promo entity.Promo

	err := r.db.First(&promo, promoId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custErr.ErrPromoNotFound
		}

		return nil, err
	}

	return &promo, nil
}

func (r *PromoRepositoryImpl) GetAllPromo(offset int, limit int, search string, orderAndSort string) ([]entity.Promo, error) {
	var promos []entity.Promo

	err := r.db.
		Where(`name ILIKE CONCAT('%',?,'%')`, search).
		Offset(offset).
		Limit(limit).
		Order(orderAndSort).
		Find(&promos).Error

	if err != nil {
		return nil, err
	}

	return promos, nil
}

func (r *PromoRepositoryImpl) CreatePromo(promo entity.Promo) (*entity.Promo, error) {
	newPromo := entity.Promo{
		Name:        promo.Name,
		MinFee:      promo.MinFee,
		Discount:    promo.Discount,
		MaxDiscount: promo.MaxDiscount,
		Quota:       promo.Quota,
		ExpireDate:  promo.ExpireDate,
	}

	if err := r.db.Omit("created_at", "updated_at", "deleted_at").Create(&newPromo).Error; err != nil {
		return nil, err
	}

	return &newPromo, nil
}

func (r *PromoRepositoryImpl) UpdatePromo(promo entity.Promo) (*entity.Promo, error) {
	newPromo := entity.Promo{
		Name:        promo.Name,
		MinFee:      promo.MinFee,
		Discount:    promo.Discount,
		MaxDiscount: promo.MaxDiscount,
		Quota:       promo.Quota,
		ExpireDate:  promo.ExpireDate,
	}
	res := r.db.
		Clauses(clause.Returning{}).
		Omit("created_at", "updated_at", "deleted_at").
		Where("id = ?", promo.Id).
		Updates(&newPromo)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, custErr.ErrPromoNotFound
	}

	return &newPromo, nil
}

func (r *PromoRepositoryImpl) CountPromo(search string) int64 {
	var totalPromo int64
	r.db.Model(&entity.Promo{}).Where(`name ILIKE CONCAT('%',?,'%')`, search).Count(&totalPromo)

	return totalPromo
}

func (r *PromoRepositoryImpl) GetAllPromoGame() ([]entity.Promo, error) {
	var promos []entity.Promo

	currentTime := time.Now()

	err := r.db.
		Where("quota > ?", 0).
		Where("expire_date > ?", currentTime.Format("2006-01-02")).
		Find(&promos).Error

	if err != nil {
		return nil, err
	}

	return promos, nil
}
