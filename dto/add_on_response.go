package dto

import "final-project-backend/entity"

type AddOnResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func CreateAddOnResponse(addOn entity.AddOn) AddOnResponse {
	return AddOnResponse{
		Id:          addOn.Id,
		Name:        addOn.Name,
		Description: addOn.Description,
		Price:       addOn.Price,
	}
}

func CreateAddOnListResponse(addOns []entity.AddOn) []AddOnResponse {
	addOnsResponse := []AddOnResponse{}
	for _, a := range addOns {
		addOn := CreateAddOnResponse(a)
		addOnsResponse = append(addOnsResponse, addOn)
	}

	return addOnsResponse
}
