package dto

type ShippingReportPaginateResponse struct {
	Page           int                `json:"page"`
	Limit          int                `json:"limit"`
	TotalPage      int                `json:"total_page"`
	Totaldata      int                `json:"total_data"`
	TotalCostMonth int                `json:"total_cost_month"`
	Data           []ShippingResponse `json:"data"`
}
