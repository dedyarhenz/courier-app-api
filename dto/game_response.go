package dto

import "time"

type GameResponse struct {
	Name        string    `json:"name"`
	MinFee      int       `json:"min_fee"`
	Discount    int       `json:"discount"`
	MaxDiscount int       `json:"max_discount"`
	ExpireDate  time.Time `json:"expire_date"`
}
