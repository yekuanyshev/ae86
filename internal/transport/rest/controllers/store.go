package controllers

import "github.com/supernova0730/ae86/internal/interfaces/container"

type StoreController struct {
	service container.IService
}

func NewStoreController(service container.IService) *StoreController {
	return &StoreController{service: service}
}
