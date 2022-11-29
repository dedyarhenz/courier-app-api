package dto

type PromoUpdateRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	MinFee      int    `json:"min_fee" binding:"required,numeric"`
	Discount    int    `json:"discount" binding:"required,numeric"`
	MaxDiscount int    `json:"max_discount" binding:"required,numeric"`
	Quota       int    `json:"quota" binding:"required,numeric"`
	ExpireDate  string `json:"expire_date" binding:"required,datetime=2006-01-02T15:04:05"`
}
