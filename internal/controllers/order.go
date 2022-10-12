package controllers

import "github.com/supernova0730/ae86/internal/interfaces/container"

type OrderController struct {
	service container.IService
}

func NewOrderController(service container.IService) *OrderController {
	return &OrderController{service: service}
}
