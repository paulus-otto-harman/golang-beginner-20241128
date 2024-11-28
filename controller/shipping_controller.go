package controller

import (
	"20241128/model"
	"20241128/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ShippingController struct {
	service service.ShippingService
}

func NewShippingController(db *gorm.DB) *ShippingController {
	return &ShippingController{
		service: service.NewShippingService(db),
	}
}

func (s *ShippingController) Cost(ctx *gin.Context) {
	destinationLat, _ := strconv.ParseFloat(ctx.Query("lat"), 32)
	destinationLng, _ := strconv.ParseFloat(ctx.Query("lng"), 32)
	courierId, _ := strconv.Atoi(ctx.Query("courier"))

	shipping := model.Shipping{CourierId: courierId, Items: 3, Destination: model.Coordinate{Lat: float32(destinationLat), Lng: float32(destinationLng)}}

	if err := s.service.GetCost(&shipping); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, shipping)
}
