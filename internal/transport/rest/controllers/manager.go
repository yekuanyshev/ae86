package controllers

import "github.com/supernova0730/ae86/internal/interfaces/container"

type ManagerController struct {
	service container.IService
}

func NewManagerController(service container.IService) *ManagerController {
	return &ManagerController{service: service}
}
