package usecase

import (
	"final-project-backend/config"
	"final-project-backend/dto"
	"final-project-backend/pkg/database/postgres"
	"final-project-backend/repository"
	"fmt"
	"testing"
)

var dbcfg = config.DatabaseConfig{
	Host:     "localhost",
	Port:     "5432",
	DbName:   "final_project",
	User:     "dedyirawan",
	Password: "dedyirawan",
}
var db = postgres.New(&config.Config{Database: dbcfg})

func TestShippingUsecaseImpl_CreateShipping_Success(t *testing.T) {
	// mockRepoShipping := mocks.NewShippingRepository(t)
	// mockRepoPayment := mocks.NewPaymentRepository(t)
	// mockRepoUser := mocks.NewUserRepository(t)
	// mockRepoAddress := mocks.NewAddressRepository(t)
	// mockRepoSize := mocks.NewSizeRepository(t)
	// mockRepoCategory := mocks.NewCategoryRepository(t)
	// mocksAddOn := mocks.NewAddOnRepository(t)

	mockRepoShipping := repository.NewShippingRepositoryImpl(db)
	mockRepoPayment := repository.NewPaymentRepositoryImpl(db)
	mockRepoUser := repository.NewUserRepositoryImpl(db)
	mockRepoAddress := repository.NewAddressRepositoryImpl(db)
	mockRepoSize := repository.NewSizeRepositoryImpl(db)
	mockRepoCategory := repository.NewCategoryRepositoryImpl(db)
	mockAddOn := repository.NewAddOnRepositoryImpl(db)

	usecase := NewShippingUsecaseImpl(mockRepoShipping, mockRepoPayment, mockRepoUser, mockRepoAddress, mockRepoSize, mockRepoCategory, mockAddOn)

	newShipping := dto.ShippingCreateRequest{
		UserId:     1,
		SizeId:     1,
		CategoryId: 3,
		AddressId:  1,
		AddOnsId:   []int{1, 2},
	}

	res, err := usecase.CreateShipping(newShipping)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
