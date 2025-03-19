package entity

import (
	"time"

	"gorm.io/gorm"
)

type Business struct {
	ID          string         `gorm:"column:id;primaryKey;"`
	Name        string         `gorm:"column:name;"`
	Industry    string         `gorm:"column:industry;"`
	CompanySize *string        `gorm:"column:company_size;"`
	Logo        *string        `gorm:"column:logo;"`
	Email       *string        `gorm:"column:email;"`
	Phone       *string        `gorm:"column:phone;"`
	Website     *string        `gorm:"column:website;"`
	Address1    *string        `gorm:"column:address_1;"`
	Address2    *string        `gorm:"column:address_2;"`
	City        *string        `gorm:"column:city;"`
	State       *string        `gorm:"column:state;"`
	Zipcode     *string        `gorm:"column:zipcode;"`
	Country     *string        `gorm:"column:country;"`
	CreatedAt   *time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Business) TableName() string {
	return "businesses"
}
