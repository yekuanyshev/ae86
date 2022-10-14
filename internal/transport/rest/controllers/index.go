package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"github.com/supernova0730/ae86/internal/transport/rest/response"
	"github.com/supernova0730/ae86/internal/transport/utils"
	"net/http"
)

type IndexController struct {
	service container.IService
}

func NewIndexController(service container.IService) *IndexController {
	return &IndexController{service: service}
}

func (ctl *IndexController) Get(c *fiber.Ctx) error {
	storeID := utils.GetMeta(c.UserContext()).StoreID
	result, err := ctl.service.Index().Get(c.UserContext(), storeID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return c.Status(http.StatusOK).JSON(result)
}
