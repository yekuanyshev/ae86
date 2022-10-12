package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"net/http"
)

type IndexController struct {
	service container.IService
}

func NewIndexController(service container.IService) *IndexController {
	return &IndexController{service: service}
}

func (ctl *IndexController) Get(c *fiber.Ctx) error {
	storeID, err := c.ParamsInt("store_id", 0)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result, err := ctl.service.Index().Get(c.Context(), int64(storeID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}
