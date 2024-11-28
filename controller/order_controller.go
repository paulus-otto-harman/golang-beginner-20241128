package controller

import (
	"20241128/model"
	"20241128/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type OrderController struct {
	service service.OrderService
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{
		service: service.NewOrderService(db),
	}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order model.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateOrder(&order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}
