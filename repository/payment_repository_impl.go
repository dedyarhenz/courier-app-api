package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepositoryImpl(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{
		db: db,
	}
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
