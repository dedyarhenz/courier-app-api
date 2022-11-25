package dto

type PromoPaginateResponse struct {
	Page      int             `json:"page"`
	Limit     int             `json:"limit"`
	TotalPage int             `json:"total_page"`
	Totaldata int             `json:"total_data"`
	Data      []PromoResponse `json:"data"`
}
