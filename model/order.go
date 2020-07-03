package model

import "github.com/jinzhu/gorm"

//Order --> Model to entity Order
type Order struct {
	gorm.Model
	UserID    uint
	ProductID uint
}
