package entity

import (
	"time"

	"gorm.io/gorm"
)

const (
	PAYMENT_PENDING = "PENDING"
	pAYMENT_SUCCESS = "SUCCESS"
)

type Payment struct {
	Id            int            `gorm:"primaryKey;column:id"`
	PaymentStatus string         `gorm:"column:payment_status"`
	TotalCost     int            `gorm:"column:total_cost"`
	PromoId       *int           `gorm:"column:promo_id"`
	CreatedAt     time.Time      `gorm:"column:created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`
}
