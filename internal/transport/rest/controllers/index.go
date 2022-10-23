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

// Get godoc
// @Summary index router
// @Description get index data
// @Produce json
// @Param X-Store-Id header int true "Store ID"
// @Success 200 {object} views.Index
// @Failure 500 {object} response.ErrorResponse
// @Router /index [get]
func (ctl *IndexController) Get(c *fiber.Ctx) error {
	storeID := utils.GetMeta(c.UserContext()).StoreID
	result, err := ctl.service.Index().Get(c.UserContext(), storeID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return c.Status(http.StatusOK).JSON(result)
}
