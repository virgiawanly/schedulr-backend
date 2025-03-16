package entity

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID         string         `gorm:"primaryKey"`
	BusinessID string         `gorm:"column:business_id"`
	Email      string         `gorm:"column:email"`
	Password   string         `gorm:"column:password"`
	Role       string         `gorm:"column:role"`
	CreatedAt  *time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  *time.Time     `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Account) TableName() string {
	return "accounts"
}
