package controllers

import "github.com/supernova0730/ae86/internal/interfaces/container"

type CategoryController struct {
	service container.IService
}

func NewCategoryController(service container.IService) *CategoryController {
	return &CategoryController{service: service}
}
