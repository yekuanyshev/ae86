package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/internal/interfaces/container"
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
	storeID := utils.GetStoreID(c.UserContext())
	result, err := ctl.service.Index().Get(c.Context(), storeID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}
