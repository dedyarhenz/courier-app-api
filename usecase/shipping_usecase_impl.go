package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"
	"final-project-backend/repository"
)

type ShippingUsecaseImpl struct {
	repoShipping repository.ShippingRepository
	repoPayment  repository.PaymentRepository
	repoUser     repository.UserRepository
	repoAddress  repository.AddressRepository
	repoSize     repository.SizeRepository
	repoCategory repository.CategoryRepository
	repoAddOn    repository.AddOnRepository
}

func NewShippingUsecaseImpl(
	repoShipping repository.ShippingRepository,
	repoPayment repository.PaymentRepository,
	repoUser repository.UserRepository,
	repoAddress repository.AddressRepository,
	repoSize repository.SizeRepository,
	repoCategory repository.CategoryRepository,
	repoAddOn repository.AddOnRepository,
) ShippingUsecase {
	return &ShippingUsecaseImpl{
		repoShipping: repoShipping,
		repoPayment:  repoPayment,
		repoUser:     repoUser,
		repoAddress:  repoAddress,
		repoSize:     repoSize,
		repoCategory: repoCategory,
		repoAddOn:    repoAddOn,
	}
}

func (u *ShippingUsecaseImpl) GetAllShippingByUserId(userId int) ([]dto.ShippingResponse, error) {
	shippings, err := u.repoShipping.GetAllShippingByUserId(userId)
	if err != nil {
		return nil, err
	}

	resShippings := dto.CreateShippingListResponse(shippings)

	return resShippings, nil
}

func (u *ShippingUsecaseImpl) GetShippingByUserId(userId int, shippingId int) (*dto.ShippingResponse, error) {
	shipping, err := u.repoShipping.GetShippingByUserId(userId, shippingId)
	if err != nil {
		return nil, err
	}

	resShippings := dto.CreateShippingResponse(*shipping)

	return &resShippings, nil
}

func (u *ShippingUsecaseImpl) CreateShipping(request dto.ShippingCreateRequest) (*dto.ShippingResponse, error) {
	if len(request.AddOnsId) == 0 {
		return nil, custErr.ErrMinAddOns
	}

	addOns, err := u.repoAddOn.GetAddOnByMultipleId(request.AddOnsId)
	if err != nil {
		return nil, nil
	}

	size, err := u.repoSize.GetSizeById(request.SizeId)
	if err != nil {
		return nil, err
	}

	category, err := u.repoCategory.GetCategoryById(request.CategoryId)
	if err != nil {
		return nil, err
	}

	_, err = u.repoUser.GetUserById(request.UserId)
	if err != nil {
		return nil, err
	}

	address, err := u.repoAddress.GetAddressBySpecificUser(request.UserId, request.AddressId)
	if err != nil {
		return nil, err
	}

	if len(addOns) != len(request.AddOnsId) {
		return nil, custErr.ErrAddOnInvalid
	}

	totalCost := calculateCost(addOns, *size, *category)
	payment, err := u.repoPayment.CreatePayment(entity.Payment{PaymentStatus: entity.PAYMENT_PENDING, TotalCost: totalCost})
	if err != nil {
		return nil, err
	}

	var newAddOnShipping []entity.AddOnShipping
	for _, addOn := range addOns {
		newAddOnShipping = append(newAddOnShipping, entity.AddOnShipping{AddOnId: addOn.Id})
	}

	newShipping := entity.Shipping{
		SizeId:         size.Id,
		CategoryId:     category.Id,
		AddressId:      address.Id,
		PaymentId:      payment.Id,
		AddOnShippings: newAddOnShipping,
		StatusShipping: entity.SHIPP_PROCESS,
	}

	shipping, err := u.repoShipping.CreateShipping(newShipping)
	if err != nil {
		return nil, err
	}

	resShipping := dto.CreateShippingResponse(*shipping)

	return &resShipping, nil
}

func calculateCost(addOns []entity.AddOn, size entity.Size, category entity.Category) int {
	totalAddOn := 0
	for _, addOn := range addOns {
		totalAddOn += addOn.Price
	}

	totalCost := totalAddOn + size.Price + category.Price

	return totalCost
}
