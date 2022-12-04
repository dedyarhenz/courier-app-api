package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	custErr "final-project-backend/pkg/errors"
	"final-project-backend/pkg/helper"
	"final-project-backend/repository"
	"fmt"
	"strings"
	"time"
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

func (u *ShippingUsecaseImpl) GetAllShipping(page int, limit int, search string, order string, sortBy string) (dto.ShippingPaginateResponse, error) {
	orderAndSort := fmt.Sprintf("%s %s", u.checkOrderShipping(order), helper.CheckSortBy(sortBy))
	offset := (page * limit) - limit
	totalData := u.repoShipping.CountShipping(search)
	totalPage := totalData / int64(limit)

	if (totalPage % int64(limit)) != 0 {
		totalPage += 1
	}

	resShippingPaginate := dto.ShippingPaginateResponse{
		Page:      page,
		Limit:     limit,
		Totaldata: int(totalData),
		TotalPage: int(totalPage),
		Data:      []dto.ShippingResponse{},
	}

	shippings, err := u.repoShipping.GetAllShipping(offset, limit, search, orderAndSort)
	if err != nil {
		return resShippingPaginate, err
	}

	resShippings := dto.CreateShippingListResponse(shippings)
	resShippingPaginate.Data = resShippings

	return resShippingPaginate, nil
}

func (u *ShippingUsecaseImpl) GetAllReportShippingByDate(month int, year int, page int, limit int, sortBy string) (dto.ShippingReportPaginateResponse, error) {
	startDate := time.Date(year, time.Month(month), 01, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
	endDate := time.Date(year, time.Month(month), 31, 0, 0, 0, 0, time.UTC).Format("2006-01-02")

	orderAndSort := fmt.Sprintf("%s %s", "created_at", helper.CheckSortBy(sortBy))
	offset := (page * limit) - limit
	totalData := u.repoShipping.CountShippingByDate(startDate, endDate)
	totalPage := totalData / int64(limit)

	if (totalPage % int64(limit)) != 0 {
		totalPage += 1
	}

	resShippingReportPaginate := dto.ShippingReportPaginateResponse{
		Page:           page,
		Limit:          limit,
		Totaldata:      int(totalData),
		TotalPage:      int(totalPage),
		TotalCostMonth: 0,
		Data:           []dto.ShippingResponse{},
	}

	shippings, err := u.repoShipping.GetAllReportShippingByDate(startDate, endDate, offset, limit, orderAndSort)
	if err != nil {
		return resShippingReportPaginate, err
	}

	totalCostMonth := u.repoPayment.TotalCostPaymentByDate(startDate, endDate)

	resShippings := dto.CreateShippingListResponse(shippings)
	resShippingReportPaginate.Data = resShippings
	resShippingReportPaginate.TotalCostMonth = int(totalCostMonth)

	return resShippingReportPaginate, nil
}

func (u *ShippingUsecaseImpl) GetShippingById(shippingId int) (*dto.ShippingResponse, error) {
	shipping, err := u.repoShipping.GetShippingById(shippingId)
	if err != nil {
		return nil, err
	}

	resShippings := dto.CreateShippingResponse(*shipping)

	return &resShippings, nil
}

func (u *ShippingUsecaseImpl) GetAllShippingByUserId(userId int, page int, limit int, search string, order string, sortBy string) (dto.ShippingPaginateResponse, error) {
	orderAndSort := fmt.Sprintf("%s %s", u.checkOrderShipping(order), helper.CheckSortBy(sortBy))
	offset := (page * limit) - limit
	totalData := u.repoShipping.CountShippingByUserId(userId, search)
	totalPage := totalData / int64(limit)

	if (totalData % int64(limit)) != 0 {
		totalPage += 1
	}

	resShippingPaginate := dto.ShippingPaginateResponse{
		Page:      page,
		Limit:     limit,
		Totaldata: int(totalData),
		TotalPage: int(totalPage),
		Data:      []dto.ShippingResponse{},
	}

	shippings, err := u.repoShipping.GetAllShippingByUserId(userId, offset, limit, search, orderAndSort)
	if err != nil {
		return resShippingPaginate, err
	}

	resShippings := dto.CreateShippingListResponse(shippings)
	resShippingPaginate.Data = resShippings

	return resShippingPaginate, nil
}

func (u *ShippingUsecaseImpl) GetShippingByUserId(userId int, shippingId int) (*dto.ShippingResponse, error) {
	shipping, err := u.repoShipping.GetShippingByUserId(userId, shippingId)
	if err != nil {
		return nil, err
	}

	resShippings := dto.CreateShippingResponse(*shipping)

	return &resShippings, nil
}

func (u *ShippingUsecaseImpl) UpdateReviewByUserId(request dto.ShippingReviewRequest) error {
	shipping, err := u.repoShipping.GetShippingByUserId(request.UserId, request.ShippingId)
	if err != nil {
		return err
	}

	if shipping.Review != nil {
		return custErr.ErrShippingAlreadyReview
	}

	if shipping.StatusShipping != entity.SHIPP_DELIVERED {
		return custErr.ErrShippingReview
	}

	if shipping.Payment.PaymentStatus != entity.PAYMENT_SUCCESS {
		return custErr.ErrShippingReview
	}

	err = u.repoShipping.UpdateReviewByUserId(request.UserId, request.ShippingId, request.Review)
	if err != nil {
		return err
	}

	return nil
}

func (u *ShippingUsecaseImpl) UpdateStatusShipping(request dto.ShippingUpdateStatusRequest) error {
	var statusShippig string

	switch strings.ToUpper(request.StatusShipping) {
	case entity.SHIPP_PROCESS:
		statusShippig = entity.SHIPP_PROCESS
	case entity.SHIPP_PICKUP:
		statusShippig = entity.SHIPP_PICKUP
	case entity.SHIPP_DELIVERY:
		statusShippig = entity.SHIPP_DELIVERY
	case entity.SHIPP_DELIVERED:
		statusShippig = entity.SHIPP_DELIVERED
	default:
		return custErr.ErrShippingStatus
	}

	shipping, err := u.repoShipping.GetShippingById(request.ShippingId)
	if err != nil {
		return err
	}

	if shipping.Payment.PaymentStatus == entity.PAYMENT_PENDING {
		return custErr.ErrShippingMustPaid
	}

	err = u.repoShipping.UpdateStatusShipping(request.ShippingId, statusShippig)
	if err != nil {
		return err
	}

	return nil
}

func (u *ShippingUsecaseImpl) CreateShipping(request dto.ShippingCreateRequest) (*dto.ShippingResponse, error) {
	if len(request.AddOnsId) == 0 {
		return nil, custErr.ErrMinAddOns
	}

	addOns, err := u.repoAddOn.GetAddOnByMultipleId(request.AddOnsId)
	if err != nil {
		return nil, err
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

	address, err := u.repoAddress.GetAddressByUserId(request.UserId, request.AddressId)
	if err != nil {
		return nil, err
	}

	if len(addOns) != len(request.AddOnsId) {
		return nil, custErr.ErrAddOnInvalid
	}

	totalCost := u.calculateCost(addOns, *size, *category)
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

func (u *ShippingUsecaseImpl) calculateCost(addOns []entity.AddOn, size entity.Size, category entity.Category) int {
	totalAddOn := 0
	for _, addOn := range addOns {
		totalAddOn += addOn.Price
	}

	totalCost := totalAddOn + size.Price + category.Price

	return totalCost
}

func (u *ShippingUsecaseImpl) checkOrderShipping(order string) string {
	switch order {
	case "date":
		order = "created_at"
	case "category":
		order = "category_id"
	case "size":
		order = "size_id"
	case "payment":
		order = "payments.total_cost"
	case "status":
		order = "status_shipping"
	default:
		order = "created_at"
	}

	return order
}
