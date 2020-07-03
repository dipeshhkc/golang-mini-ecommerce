package handler

import (
	"mini-ecommerce/model"
	"mini-ecommerce/repository"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)


//ProductHandler --> interface to Product handler 
type ProductHandler interface {
	AddProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
}

type productHandler struct{
	repo repository.ProductRepository
} 

//NewProductHandler --> returns new handler for product entity
func NewProductHandler() ProductHandler {
	return &productHandler{
		repo: repository.NewProductRepository(),
	}
}

func(h *productHandler) AddProduct(ctx *gin.Context){
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	product, err := h.repo.AddProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}
	ctx.JSON(http.StatusOK, product)

}
func(h *productHandler) UpdateProduct(ctx *gin.Context){

	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	product.ID = uint(intID)
	product, err = h.repo.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, product)

}
func(h *productHandler) DeleteProduct(ctx *gin.Context){

	var product model.Product
	id := ctx.Param("id")
	intID, _ := strconv.Atoi(id)
	product.ID = uint(intID)
	product, err := h.repo.DeleteProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, product)

}