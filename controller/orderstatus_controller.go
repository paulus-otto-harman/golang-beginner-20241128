package controller

import (
	"20241128/model"
	"20241128/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type OrderStatusController struct {
	service service.OrderStatusService
}

func NewOrderStatusController(db *gorm.DB) *OrderStatusController {
	return &OrderStatusController{
		service: service.NewOrderStatusService(db),
	}
}

func (c *OrderStatusController) CreateOrderStatus(ctx *gin.Context) {
	orderId, _ := strconv.Atoi(ctx.Param("id"))
	orderStatus := model.OrderStatus{OrderID: uint(orderId)}
	if err := ctx.ShouldBindJSON(&orderStatus); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateOrderStatus(&orderStatus); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, orderStatus)
}

func (c *OrderStatusController) GetByOrderId(ctx *gin.Context) {
	orderId, _ := strconv.Atoi(ctx.Param("id"))
	orderStatus, err := c.service.GetByOrderId(uint(orderId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, orderStatus)
}
