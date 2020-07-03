package repository

import (
	"mini-ecommerce/model"

	"github.com/jinzhu/gorm"
)

//ProductRepository --> Interface to ProductRepository
type ProductRepository interface {
	AddProduct(model.Product)  (model.Product,error)
	UpdateProduct(model.Product) (model.Product,error)
	DeleteProduct(model.Product)  (model.Product,error)
}

type productRepository struct {
	connection *gorm.DB
}

//NewProductRepository --> returns new product repository
func NewProductRepository() ProductRepository {
	return &productRepository{
		connection: DB(),
	}
}

func(db *productRepository) AddProduct(product model.Product)  (model.Product,error){
	return product, db.connection.Create(&product).Error
}

func(db *productRepository) UpdateProduct(product model.Product) (model.Product,error){
	if err := db.connection.First(&product, product.ID).Error; err != nil {
		return product, err
	}
	return product, db.connection.Model(&product).Updates(&product).Error
}

func(db *productRepository) DeleteProduct(product model.Product)  (model.Product,error){
	if err := db.connection.First(&product, product.ID).Error; err != nil {
		return product, err
	}
	return product, db.connection.Delete(&product).Error
}
