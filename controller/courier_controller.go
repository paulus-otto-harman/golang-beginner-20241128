package controller

import (
	"20241128/model"
	"20241128/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type CourierController struct {
	service service.CourierService
}

func NewCourierController(db *gorm.DB) *CourierController {
	return &CourierController{
		service: service.NewCourierService(db),
	}
}

func (c *CourierController) GetCouriers(ctx *gin.Context) {
	couriers, err := c.service.GetAllCouriers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, couriers)
}

func (c *CourierController) CreateCourier(ctx *gin.Context) {
	var courier model.Courier
	if err := ctx.ShouldBindJSON(&courier); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateCourier(&courier); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, courier)
}
