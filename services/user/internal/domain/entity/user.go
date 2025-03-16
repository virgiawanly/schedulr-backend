package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         string         `gorm:"primaryKey"`
	BusinessID string         `gorm:"column:business_id"`
	AccountID  string         `gorm:"column:account_id"`
	FirstName  string         `gorm:"column:first_name"`
	LastName   string         `gorm:"column:last_name"`
	Phone      string         `gorm:"column:phone"`
	Address1   string         `gorm:"column:address_1"`
	Address2   string         `gorm:"column:address_2"`
	City       string         `gorm:"column:city"`
	State      string         `gorm:"column:state"`
	Zipcode    string         `gorm:"column:zipcode"`
	Country    string         `gorm:"column:country"`
	LabourCost float64        `gorm:"column:labour_cost"`
	Language   string         `gorm:"column:language"`
	Image      string         `gorm:"column:image"`
	CreatedAt  *time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  *time.Time     `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (User) TableName() string {
	return "users"
}
