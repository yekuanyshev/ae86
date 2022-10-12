package controllers

import "github.com/supernova0730/ae86/internal/interfaces/container"

type BannerController struct {
	service container.IService
}

func NewBannerController(service container.IService) *BannerController {
	return &BannerController{service: service}
}
