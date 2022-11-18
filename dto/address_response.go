package dto

import "final-project-backend/entity"

type AddressResponse struct {
	Id             int    `json:"id"`
	RecipientName  string `json:"recipient_name"`
	FullAddress    string `json:"full_address"`
	RecipientPhone string `json:"recipient_phone"`
	UserId         int    `json:"user_id"`
}

func CreateAddressResponse(address entity.Address) AddressResponse {
	return AddressResponse{
		Id:             address.Id,
		RecipientName:  address.RecipientName,
		FullAddress:    address.FullAddress,
		RecipientPhone: address.RecipientPhone,
		UserId:         address.UserId,
	}
}

func CreateAddressListResponse(addresses []entity.Address) []AddressResponse {
	addressesResponse := []AddressResponse{}
	for _, a := range addresses {
		address := CreateAddressResponse(a)
		addressesResponse = append(addressesResponse, address)
	}

	return addressesResponse
}
