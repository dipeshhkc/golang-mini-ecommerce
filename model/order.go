package model

import (
	"github.com/jinzhu/gorm"
)

//Order --> Model to entity Order
type Order struct {
	gorm.Model
	User      User    `gorm:"foreignkey:UserID"`
	Product   Product `gorm:"foreignkey:ProductID"`
	UserID    uint
	ProductID uint
	Quantity  int `json:"quantity"`
}

//TableName --> Table for Order Model
func (Order) TableName() string {
	return "orders"
}
