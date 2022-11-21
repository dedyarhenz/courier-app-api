package repository

import "final-project-backend/entity"

type PaymentRepository interface {
	CreatePayment(payment entity.Payment) (*entity.Payment, error)
}
