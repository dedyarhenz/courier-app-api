package dto

type ShippingPaginateResponse struct {
	Page      int                `json:"page"`
	Limit     int                `json:"limit"`
	TotalPage int                `json:"total_page"`
	Totaldata int                `json:"total_data"`
	Data      []ShippingResponse `json:"data"`
}
