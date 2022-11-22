package entity

import (
	"time"

	"gorm.io/gorm"
)

type AddOnShipping struct {
	Id         int            `gorm:"primaryKey;column:id"`
	ShippingId int            `gorm:"column:shipping_id"`
	AddOnId    int            `gorm:"column:add_on_id"`
	Shipping   *Shipping      `gorm:"foreignKey:ShippingId;references:Id"`
	AddOn      *AddOn         `gorm:"foreignKey:AddOnId;references:Id"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`
}
