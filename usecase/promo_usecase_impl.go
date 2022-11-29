package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/pkg/helper"
	"final-project-backend/repository"
	"fmt"
	"time"
)

type PromoUsecaseImpl struct {
	repoPromo repository.PromoRepository
}

func NewPromoUsecaseImpl(repoPromo repository.PromoRepository) PromoUsecase {
	return &PromoUsecaseImpl{
		repoPromo: repoPromo,
	}
}

func (u *PromoUsecaseImpl) GetAllPromo(page int, limit int, search string, order string, sort string) (dto.PromoPaginateResponse, error) {
	orderAndSort := fmt.Sprintf("%s %s", checkOrderPromo(order), helper.CheckSortBy(sort))
	offset := (page * limit) - limit
	totalData := u.repoPromo.CountPromo(search)
	totalPage := totalData / int64(limit)

	if (totalData % int64(limit)) != 0 {
		totalPage += 1
	}

	resPromoPaginate := dto.PromoPaginateResponse{
		Page:      page,
		Limit:     limit,
		Totaldata: int(totalData),
		TotalPage: int(totalPage),
		Data:      []dto.PromoResponse{},
	}

	allPromo, err := u.repoPromo.GetAllPromo(offset, limit, search, orderAndSort)
	if err != nil {
		return resPromoPaginate, err
	}

	resAllPromo := dto.CreatePromoListResponse(allPromo)
	resPromoPaginate.Data = resAllPromo

	return resPromoPaginate, nil
}

func (u *PromoUsecaseImpl) GetPromoById(promoId int) (*dto.PromoResponse, error) {
	promo, err := u.repoPromo.GetPromoById(promoId)
	if err != nil {
		return nil, err
	}

	resPromo := dto.CreatePromoResponse(*promo)

	return &resPromo, nil
}

func (u *PromoUsecaseImpl) CreatePromo(request dto.PromoCreateRequest) (*dto.PromoResponse, error) {
	expireDate, err := time.Parse("2006-01-02T15:04:05", request.ExpireDate)
	if err != nil {
		return nil, err
	}

	newPromo := entity.Promo{
		Name:        request.Name,
		MinFee:      request.MinFee,
		Discount:    request.Discount,
		MaxDiscount: request.MaxDiscount,
		Quota:       request.Quota,
		ExpireDate:  expireDate,
	}

	promo, err := u.repoPromo.CreatePromo(newPromo)
	if err != nil {
		return nil, err
	}

	resPromo := dto.CreatePromoResponse(*promo)

	return &resPromo, nil
}

func (u *PromoUsecaseImpl) UpdatePromo(request dto.PromoUpdateRequest) (*dto.PromoResponse, error) {
	expireDate, err := time.Parse("2006-01-02", request.ExpireDate)
	if err != nil {
		return nil, err
	}

	newPromo := entity.Promo{
		Id:          request.Id,
		Name:        request.Name,
		MinFee:      request.MinFee,
		Discount:    request.Discount,
		MaxDiscount: request.MaxDiscount,
		Quota:       request.Quota,
		ExpireDate:  expireDate,
	}

	promo, err := u.repoPromo.UpdatePromo(newPromo)
	if err != nil {
		return nil, err
	}

	resPromo := dto.CreatePromoResponse(*promo)

	return &resPromo, nil
}

func checkOrderPromo(order string) string {
	switch order {
	case "expired":
		order = "expire_date"
	case "quota":
		order = "quota"
	default:
		order = "expired"
	}

	return order
}
