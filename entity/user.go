package entity

import (
	"time"

	"gorm.io/gorm"
)

const (
	UserRole  = "USER"
	AdminRole = "ADMIN"
)

type User struct {
	Id           int            `gorm:"primaryKey;column:id"`
	Email        string         `gorm:"column:email"`
	Password     string         `gorm:"column:password"`
	FullName     string         `gorm:"column:long_name"`
	Phone        string         `gorm:"column:phone"`
	Role         string         `gorm:"column:role"`
	Balance      int            `gorm:"column:balance"`
	Photo        string         `gorm:"column:photo"`
	RefferalCode string         `gorm:"column:refferal_code"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
}
