package entity

import (
	"time"

	"gorm.io/gorm"
)

type AddOn struct {
	Id          int            `gorm:"primaryKey;column:id"`
	Name        string         `gorm:"column:name"`
	Description string         `gorm:"column:description"`
	Price       int            `gorm:"column:price"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}
