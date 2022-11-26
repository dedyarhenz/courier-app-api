package usecase

import (
	"final-project-backend/dto"
	"final-project-backend/entity"
	"final-project-backend/repository"
)

type AddressUsecaseImpl struct {
	repoAddress repository.AddressRepository
}

func NewAddressUsecaseImpl(repoAddress repository.AddressRepository) AddressUsecase {
	return &AddressUsecaseImpl{
		repoAddress: repoAddress,
	}
}

func (u *AddressUsecaseImpl) GetAllAddress(page int, limit int, search string) (dto.AddressPaginateResponse, error) {
	orderAndSort := "created_at desc"
	offset := (page * limit) - limit
	totalData := u.repoAddress.CountAddress(search)
	totalPage := totalData/int64(limit) + 1

	resAddressPaginate := dto.AddressPaginateResponse{
		Page:      page,
		Limit:     limit,
		Totaldata: int(totalData),
		TotalPage: int(totalPage),
		Data:      []dto.AddressResponse{},
	}

	allAddress, err := u.repoAddress.GetAllAddress(offset, limit, search, orderAndSort)
	if err != nil {
		return resAddressPaginate, err
	}

	resAllAddress := dto.CreateAddressListResponse(allAddress)
	resAddressPaginate.Data = resAllAddress

	return resAddressPaginate, nil
}

func (u *AddressUsecaseImpl) GetAllAddressByUserId(userId int, page int, limit int, search string) (dto.AddressPaginateResponse, error) {
	orderAndSort := "created_at desc"
	offset := (page * limit) - limit
	totalData := u.repoAddress.CountAddressByUserId(userId, search)
	totalPage := totalData/int64(limit) + 1

	resAddressPaginate := dto.AddressPaginateResponse{
		Page:      page,
		Limit:     limit,
		Totaldata: int(totalData),
		TotalPage: int(totalPage),
		Data:      []dto.AddressResponse{},
	}

	userAddress, err := u.repoAddress.GetAllAddressByUserId(userId, offset, limit, search, orderAndSort)
	if err != nil {
		return resAddressPaginate, err
	}

	resUserAddress := dto.CreateAddressListResponse(userAddress)
	resAddressPaginate.Data = resUserAddress

	return resAddressPaginate, nil
}

func (u *AddressUsecaseImpl) GetAddressByUserId(userId int, addressId int) (*dto.AddressResponse, error) {
	userAddress, err := u.repoAddress.GetAddressByUserId(userId, addressId)
	if err != nil {
		return nil, err
	}

	resUserAddress := dto.CreateAddressResponse(*userAddress)

	return &resUserAddress, nil
}

func (u *AddressUsecaseImpl) CreateAddress(request dto.AddressCreateRequest) (*dto.AddressResponse, error) {
	newAddress := entity.Address{
		RecipientName:  request.RecipientName,
		FullAddress:    request.FullAddress,
		RecipientPhone: request.RecipientPhone,
		UserId:         request.UserId,
	}

	address, err := u.repoAddress.CreateAddress(newAddress)
	if err != nil {
		return nil, err
	}

	resAddress := dto.CreateAddressResponse(*address)

	return &resAddress, nil
}

func (u *AddressUsecaseImpl) UpdateAddressByUserId(request dto.AddressUpdateRequest) (*dto.AddressResponse, error) {
	newAddress := entity.Address{
		Id:             request.Id,
		RecipientName:  request.RecipientName,
		FullAddress:    request.FullAddress,
		RecipientPhone: request.RecipientPhone,
		UserId:         request.UserId,
	}

	address, err := u.repoAddress.UpdateAddressByUserId(newAddress)
	if err != nil {
		return nil, err
	}

	resAddress := dto.CreateAddressResponse(*address)

	return &resAddress, nil
}

func (u *AddressUsecaseImpl) DeleteAddressByUserId(userId int, addressId int) error {
	err := u.repoAddress.DeleteAddressByUserId(userId, addressId)
	if err != nil {
		return err
	}

	return nil
}
