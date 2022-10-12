package controllers

import "github.com/supernova0730/ae86/internal/interfaces/container"

type OrderItemController struct {
	service container.IService
}

func NewOrderItemController(service container.IService) *OrderItemController {
	return &OrderItemController{service: service}
}
