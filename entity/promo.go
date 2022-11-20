package entity

import (
	"time"

	"gorm.io/gorm"
)

type Promo struct {
	Id          int            `gorm:"primaryKey;column:id"`
	Name        string         `gorm:"column:name"`
	MinFee      int            `gorm:"column:min_fee"`
	Discount    int            `gorm:"column:discount"`
	MaxDiscount int            `gorm:"column:max_discount"`
	Quota       int            `gorm:"column:quota"`
	ExpireDate  time.Time      `gorm:"column:expire_date"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}
