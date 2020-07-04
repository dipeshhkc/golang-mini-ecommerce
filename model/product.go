package model

import "github.com/jinzhu/gorm"

//Product --> Model for Product table
type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}

//TableName --> Table for Product Model
func (Product) TableName() string {
	return "products"
}
