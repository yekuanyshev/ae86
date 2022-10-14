package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/internal/container"
	"github.com/supernova0730/ae86/internal/transport/rest/controllers"
	"github.com/supernova0730/ae86/internal/transport/rest/middlewares"
)

func RegisterRoutes(r fiber.Router) {
	controller := controllers.NewContainer(container.MContainer.Services())

	r.Use(middlewares.SetContextHolder())
	r.Use(middlewares.SetMeta())

	v1 := r.Group("/api/v1")
	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON("pong")
	})

	v1.Get("/index", controller.Index().Get)

	productRouter := v1.Group("/product")
	productRouter.Get("/search", controller.Product().Search)
	productRouter.Get("/by_id/:id", controller.Product().ByID)
	productRouter.Get("/by_category_id/:category_id", controller.Product().ListByCategoryID)

	// todo: serving static files
	fileRouter := v1.Group("/file")
	fileRouter.Post("/upload", controller.File().Upload)
	fileRouter.Post("/uploads", controller.File().Uploads)
}
