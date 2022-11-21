package repository

import (
	"final-project-backend/entity"
	"fmt"
	"testing"
)

func TestPaymentRepositoryImpl_CreatePayment(t *testing.T) {
	r := NewPaymentRepositoryImpl(db)

	var promoId *int = nil
	newPayment := entity.Payment{
		PaymentStatus: "tes status",
		TotalCost:     80000,
		PromoId:       promoId,
	}

	res, err := r.CreatePayment(newPayment)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
