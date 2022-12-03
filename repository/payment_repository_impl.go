package repository

import (
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepositoryImpl(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{
		db: db,
	}
}

func (r *PaymentRepositoryImpl) GetPaymentById(paymentId int) (*entity.Payment, error) {
	var payment entity.Payment

	err := r.db.Preload("Shipping.Address").First(&payment, paymentId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custErr.ErrPaymentNotFound
		}

		return nil, err
	}

	return &payment, nil
}

func (r *PaymentRepositoryImpl) CreatePayment(payment entity.Payment) (*entity.Payment, error) {
	newPayment := entity.Payment{
		PaymentStatus: payment.PaymentStatus,
		TotalCost:     payment.TotalCost,
		PromoId:       payment.PromoId,
	}

	err := r.db.Omit("created_at", "updated_at", "deleted_at").Create(&newPayment).Error
	if err != nil {
		return nil, err
	}

	return &newPayment, nil
}

func (r *PaymentRepositoryImpl) UpdatePayment(payment entity.Payment) (*entity.Payment, error) {
	newPayment := entity.Payment{}

	res := r.db.
		Model(&newPayment).
		Clauses(clause.Returning{}).
		Omit("created_at", "updated_at", "deleted_at").
		Where("id = ?", payment.Id).
		Updates(payment)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, custErr.ErrPaymentNotFound
	}

	return &newPayment, nil
}

func (r *PaymentRepositoryImpl) TotalCostPaymentSuccessUser(userId int) int64 {
	var total int64

	r.db.
		Table("payments").
		Select("sum(total_cost)").
		Joins("INNER JOIN shippings ON shippings.payment_id = payments.id").
		Joins("INNER JOIN addresses ON addresses.id = shippings.address_id AND addresses.user_id = ?", userId).
		Where("payments.payment_status", "SUCCESS").
		Scan(&total)

	return total
}

func (r *PaymentRepositoryImpl) TotalCostPaymentByDate(startDate string, endDate string) int64 {
	var total int64

	r.db.
		Table("payments").
		Select("sum(total_cost)").
		Where("payments.payment_status", "SUCCESS").
		Where("created_at >= ?", startDate).
		Where("created_at <= ?", endDate).
		Scan(&total)

	return total
}
