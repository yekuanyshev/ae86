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

// ByID godoc
// @Summary get product
// @Description get product by id
// @Tags product
// @Produce json
// @Param X-Store-Id header int true "Store ID"
// @Param id path int true "id"
// @Success 200 {object} views.ProductFull
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /product/by_id/:id [get]
func (ctl *ProductController) ByID(c *fiber.Ctx) error {
	storeID := utils.GetMeta(c.UserContext()).StoreID

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	result, err := ctl.service.Product().ByID(c.UserContext(), storeID, int64(id))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return c.Status(http.StatusOK).JSON(result)
}

// ListByCategoryID godoc
// @Summary get products of category
// @Description get products by category id
// @Tags product
// @Produce json
// @Param X-Store-Id header int true "Store ID"
// @Param category_id path int true "category_id"
// @Success 200 {object} []views.ProductShort
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /product/by_category_id/:category_id [get]
func (ctl *ProductController) ListByCategoryID(c *fiber.Ctx) error {
	storeID := utils.GetMeta(c.UserContext()).StoreID

	categoryID, err := c.ParamsInt("category_id", 0)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	result, err := ctl.service.Product().ListByCategoryID(c.UserContext(), storeID, int64(categoryID))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return c.Status(http.StatusOK).JSON(result)
}

// Search godoc
// @Summary search products
// @Description search products by text
// @Tags product
// @Produce json
// @Param X-Store-Id header int true "Store ID"
// @Param text query string true "text"
// @Success 200 {object} views.ProductSearchResult
// @Failure 500 {object} response.ErrorResponse
// @Router /product/search [get]
func (ctl *ProductController) Search(c *fiber.Ctx) error {
	storeID := utils.GetMeta(c.UserContext()).StoreID
	searchText := c.Query("text")
	result, err := ctl.service.Product().Search(c.UserContext(), storeID, searchText)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	return c.Status(http.StatusOK).JSON(result)
}
