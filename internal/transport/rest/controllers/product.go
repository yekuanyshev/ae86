package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"github.com/supernova0730/ae86/internal/transport/rest/response"
	"github.com/supernova0730/ae86/internal/transport/utils"
	"net/http"
)

type ProductController struct {
	service container.IService
}

func NewProductController(service container.IService) *ProductController {
	return &ProductController{service: service}
}

func (ctl *ProductController) ByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	result, err := ctl.service.Product().ByID(c.UserContext(), int64(id))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (ctl *ProductController) ListByCategoryID(c *fiber.Ctx) error {
	categoryID, err := c.ParamsInt("category_id", 0)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	result, err := ctl.service.Product().ListByCategoryID(c.UserContext(), int64(categoryID))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (ctl *ProductController) Search(c *fiber.Ctx) error {
	storeID := utils.GetMeta(c.UserContext()).StoreID
	searchText := c.Query("text")
	result, err := ctl.service.Product().Search(c.UserContext(), storeID, searchText)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return c.Status(http.StatusOK).JSON(result)
}
