package model

import "github.com/jinzhu/gorm"

//User -> model for users table
type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

//TableName --> Table for Product Model
func (User) TableName() string {
	return "users"
}
