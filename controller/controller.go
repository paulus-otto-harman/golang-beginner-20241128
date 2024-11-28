package controller

import "gorm.io/gorm"

type Controller struct {
	Courier     CourierController
	Customer    CustomerController
	Order       OrderController
	OrderStatus OrderStatusController
	Product     ProductController
	Shipping    ShippingController
}

func NewController(db *gorm.DB) Controller {
	return Controller{
		Courier:     *NewCourierController(db),
		Customer:    *NewCustomerController(db),
		Order:       *NewOrderController(db),
		OrderStatus: *NewOrderStatusController(db),
		Product:     *NewProductController(db),
		Shipping:    *NewShippingController(db),
	}
}
