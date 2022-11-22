package dto

import (
	"final-project-backend/entity"
)

type ShippingResponse struct {
	Id             int               `json:"id"`
	SizeId         int               `json:"size_id"`
	CategoryId     int               `json:"category_id"`
	AddressId      int               `json:"address_id"`
	StatusShipping string            `json:"status_shipping"`
	Review         *string           `json:"review"`
	Size           *SizeResponse     `json:"size"`
	Category       *CategoryResponse `json:"category"`
	Address        *AddressResponse  `json:"addres"`
	Payment        *PaymentResponse  `json:"payment"`
}

func CreateShippingResponse(shipping entity.Shipping) ShippingResponse {
	var size *SizeResponse
	if shipping.Size != nil {
		res := CreateSizeResponse(*shipping.Size)
		size = &res
	}

	var category *CategoryResponse
	if shipping.Category != nil {
		res := CreateCategoryResponse(*shipping.Category)
		category = &res
	}

	var address *AddressResponse
	if shipping.Address != nil {
		res := CreateAddressResponse(*shipping.Address)
		address = &res
	}

	var payment *PaymentResponse
	if shipping.Payment != nil {
		res := CreatePaymentResponse(*shipping.Payment)
		payment = &res
	}

	return ShippingResponse{
		Id:             shipping.Id,
		SizeId:         shipping.SizeId,
		CategoryId:     shipping.CategoryId,
		AddressId:      shipping.CategoryId,
		StatusShipping: shipping.StatusShipping,
		Review:         shipping.Review,
		Size:           size,
		Category:       category,
		Address:        address,
		Payment:        payment,
	}
}

func CreateShippingListResponse(shippings []entity.Shipping) []ShippingResponse {
	shippingResponse := []ShippingResponse{}
	for _, c := range shippings {
		shipping := CreateShippingResponse(c)
		shippingResponse = append(shippingResponse, shipping)
	}

	return shippingResponse
}
