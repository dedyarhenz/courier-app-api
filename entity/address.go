package entity

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	Id             int            `gorm:"primaryKey;column:id"`
	RecipientName  string         `gorm:"column:recipient_name"`
	FullAddress    string         `gorm:"column:full_address"`
	RecipientPhone string         `gorm:"column:recipient_phone"`
	UserId         int            `gorm:"column:user_id"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`
}
