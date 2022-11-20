package dto

import "final-project-backend/entity"

type SizeResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func CreateSizeResponse(size entity.Size) SizeResponse {
	return SizeResponse{
		Id:          size.Id,
		Name:        size.Name,
		Description: size.Description,
		Price:       size.Price,
	}
}

func CreateSizeListResponse(sizes []entity.Size) []SizeResponse {
	sizeResponse := []SizeResponse{}
	for _, s := range sizes {
		size := CreateSizeResponse(s)
		sizeResponse = append(sizeResponse, size)
	}

	return sizeResponse
}
