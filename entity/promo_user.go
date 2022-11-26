package entity

import (
	"time"

	"gorm.io/gorm"
)

type PromoUser struct {
	Id        int            `gorm:"primaryKey;column:id"`
	PromoId   int            `gorm:"primaryKey;column:promo_id"`
	UserId    int            `gorm:"primaryKey;column:user_id"`
	IsUsed    bool           `gorm:"primaryKey;column:is_used"`
	Promo     *Promo         `gorm:"foreignKey:PromoId;references:Id"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
