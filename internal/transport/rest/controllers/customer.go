package controllers

import "github.com/supernova0730/ae86/internal/interfaces/container"

type CustomerController struct {
	service container.IService
}

func NewCustomerController(service container.IService) *CustomerController {
	return &CustomerController{service: service}
}
