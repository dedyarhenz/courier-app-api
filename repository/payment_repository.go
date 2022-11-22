package repository

import "final-project-backend/entity"

type PaymentRepository interface {
	GetPaymentById(paymentId int) (*entity.Payment, error)
	CreatePayment(payment entity.Payment) (*entity.Payment, error)
	UpdatePayment(payment entity.Payment) (*entity.Payment, error)
}
