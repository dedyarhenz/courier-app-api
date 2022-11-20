package dto

import "final-project-backend/entity"

type CategoryResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func CreateCategoryResponse(category entity.Category) CategoryResponse {
	return CategoryResponse{
		Id:          category.Id,
		Name:        category.Name,
		Description: category.Description,
		Price:       category.Price,
	}
}

func CreateCategoryListResponse(categories []entity.Category) []CategoryResponse {
	categoryResponse := []CategoryResponse{}
	for _, c := range categories {
		category := CreateCategoryResponse(c)
		categoryResponse = append(categoryResponse, category)
	}

	return categoryResponse
}
