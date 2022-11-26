package repository

import (
	"final-project-backend/entity"
	"fmt"
	"testing"
)

func TestPromoUserRepositoryImpl_UpdatePromoUser(t *testing.T) {
	r := NewPromoUserRepositoryImpl(db)

	data := entity.PromoUser{
		Id:      1,
		PromoId: 1,
		UserId:  3,
		IsUsed:  true,
	}
	_, err := r.UpdatePromoUser(data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("s")
}
