package handler

import (
	"mini-ecommerce/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//OrderHandler --> Handler for Order Entity
type OrderHandler interface {
	OrderProduct(*gin.Context)
}

type orderHandler struct {
	repo repository.OrderRepository
}

//NewOrderHandler --> return new Order Handler
func NewOrderHandler() OrderHandler {
	return &orderHandler{
		repo: repository.NewOrderRepository(),
	}
}

func (h *orderHandler) OrderProduct(ctx *gin.Context) {
	prodIDStr := ctx.Param("product")
	if prodID, err := strconv.Atoi(prodIDStr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		quantityIDStr := ctx.Param("quantity")
		if quantityID, err := strconv.Atoi(quantityIDStr); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			userID := ctx.GetFloat64("userID")
			if err := h.repo.OrderProduct(int(userID), prodID, quantityID); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.String(http.StatusOK, "Product Successfully ordered")
			}
		}
	}

}
