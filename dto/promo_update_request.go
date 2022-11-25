package dto

type PromoUpdateRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	MinFee      int    `json:"min_fee"`
	Discount    int    `json:"discount"`
	MaxDiscount int    `json:"max_discount"`
	Quota       int    `json:"quota"`
	ExpireDate  string `json:"expire_date"`
}
