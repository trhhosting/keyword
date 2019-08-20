package gorm

import (
	"time"
)

type CustomerDataPlatforms struct {
    ID        uint `gorm:"primary_key;auto_increment",json:"id"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
    // Table custom fields
	PlatformCode            int `gorm:"auto_increment;unique",json:"platform_code"`
	CustomerPlatformDetails string `gorm:"size:100",json:"customer_platform_details"`
	CustomerName string `gorm:"association_foreignkey:CustomerName"`
}

type Customers struct {
    ID        uint `gorm:"AUTO_INCREMENT;primary_key",json:"id"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
    // Table custom fields
	CustomerName string `json:"customer_name;unique"`
	Title string `json:"title"`
	Gender string `gorm:"size:1",json:"gender"`

}

