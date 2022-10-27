package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/internal/dto"
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"github.com/supernova0730/ae86/internal/transport/rest/response"
	"github.com/supernova0730/ae86/internal/transport/utils"
	"github.com/supernova0730/ae86/internal/views"
)

type OrderController struct {
	service container.IService
}

func NewOrderController(service container.IService) *OrderController {
	return &OrderController{service: service}
}

// Create godoc
// @Summary create order
// @Description create order
// @Tags order
// @Produce json
// @Param X-Store-Id header int true "Store ID"
// @Param X-Customer-Id header int true "Customer ID"
// @Param order body dto.OrderCreateDTO true "body"
// @Success 200 {object} views.OrderCreate
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /order [post]
func (ctl *OrderController) Create(c *fiber.Ctx) error {
	var orderCreateDTO dto.OrderCreateDTO

	if err := c.BodyParser(&orderCreateDTO); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	customerID := utils.GetCustomerID(c.UserContext())
	storeID := utils.GetMeta(c.UserContext()).StoreID

	id, err := ctl.service.Order().Create(c.UserContext(), customerID, storeID, orderCreateDTO)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return c.Status(http.StatusOK).JSON(views.OrderCreate{
		ID: id,
	})
}
