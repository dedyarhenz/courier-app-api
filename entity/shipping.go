package entity

import (
	"time"

	"gorm.io/gorm"
)

const (
	SHIPP_PROCESS   = "PROCESS"
	SHIPP_PICKUP    = "PICKUP"
	SHIPP_DELIVERY  = "DELIVERY"
	SHIPP_DELIVERED = "DELIVERED"
)

type Shipping struct {
	Id             int             `gorm:"primaryKey;column:id"`
	SizeId         int             `gorm:"column:size_id"`
	CategoryId     int             `gorm:"column:category_id"`
	AddressId      int             `gorm:"column:address_id"`
	PaymentId      int             `gorm:"column:payment_id"`
	StatusShipping string          `gorm:"column:status_shipping"`
	Review         *string         `gorm:"column:review"`
	IsPlayGame     bool            `gorm:"column:is_play_game"`
	Size           *Size           `gorm:"foreignKey:SizeId;references:Id"`
	Category       *Category       `gorm:"foreignKey:CategoryId;references:Id"`
	Address        *Address        `gorm:"foreignKey:AddressId;references:Id"`
	Payment        *Payment        `gorm:"foreignKey:PaymentId;references:Id"`
	AddOnShippings []AddOnShipping `gorm:"foreignKey:ShippingId;references:Id"`
	AddOns         []AddOn         `gorm:"many2many:add_on_shippings;foreignKey:Id;joinForeignKey:ShippingId;References:Id;joinReferences:AddOnId"`
	CreatedAt      time.Time       `gorm:"column:created_at"`
	UpdatedAt      time.Time       `gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt  `gorm:"column:deleted_at"`
}
