package repository

import (
	"final-project-backend/entity"
	"fmt"
	"testing"
)

func TestShippingRepositoryImpl_CreateShipping(t *testing.T) {
	newShipping := entity.Shipping{
		SizeId:         1,
		CategoryId:     3,
		AddressId:      1,
		PaymentId:      2,
		StatusShipping: "tes asdas",
		AddOnShippings: []entity.AddOnShipping{
			{
				AddOnId: 1,
			},
		},
	}

	r := NewShippingRepositoryImpl(db)
	res, err := r.CreateShipping(newShipping)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestShippingRepositoryImpl_GetShippingByUserId(t *testing.T) {
	r := NewShippingRepositoryImpl(db)
	res, err := r.GetAllShippingByUserId(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", res)
}
