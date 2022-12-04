package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromoUserUsecaseImpl_GetAllPromoUserByUserId_Error(t *testing.T) {
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)
	usecase := NewPromoUserUsecaseImpl(mockPromoUserRepo)

	mockPromoUserRepo.On("GetAllPromoUserByUserId", 1).Return(nil, fmt.Errorf("error"))
	res, err := usecase.GetAllPromoUserByUserId(1)

	assert.Nil(t, res)
	assert.Equal(t, fmt.Errorf("error"), err)
}

func TestPromoUserUsecaseImpl_GetAllPromoUserByUserId_Success(t *testing.T) {
	mockPromoUserRepo := mocks.NewPromoUserRepository(t)
	usecase := NewPromoUserUsecaseImpl(mockPromoUserRepo)

	promoUsers := []entity.PromoUser{
		{
			Id:      1,
			PromoId: 1,
			UserId:  1,
			IsUsed:  false,
		},
		{
			Id:      2,
			PromoId: 2,
			UserId:  1,
			IsUsed:  false,
		},
	}

	promoUsersResponse := dto.CreatePromoUserListResponse(promoUsers)

	mockPromoUserRepo.On("GetAllPromoUserByUserId", 1).Return(promoUsers, nil)
	res, err := usecase.GetAllPromoUserByUserId(1)

	assert.Nil(t, err)
	assert.Equal(t, promoUsersResponse, res)
}
