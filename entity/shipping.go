package entity

import (
	"time"

	"gorm.io/gorm"
)

type Shipping struct {
	Id             int            `gorm:"primaryKey;column:id"`
	SizeId         int            `gorm:"column:size_id"`
	CategoryId     int            `gorm:"column:category_id"`
	AddressId      int            `gorm:"column:address_id"`
	PaymentId      int            `gorm:"column:payment_id"`
	StatusShipping string         `gorm:"column:status_shipping"`
	Review         *string        `gorm:"column:review"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`
}
