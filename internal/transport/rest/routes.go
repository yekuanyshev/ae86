package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/supernova0730/ae86/internal/container"
	"github.com/supernova0730/ae86/internal/transport/rest/controllers"
	"github.com/supernova0730/ae86/internal/transport/rest/middlewares"
)

func RegisterRoutes(r fiber.Router) {
	controller := controllers.NewContainer(container.MContainer.Services())

	r.Use(middlewares.Recover())
	r.Use(cors.New())
	r.Use(middlewares.SetContextHolder())
	r.Use(middlewares.SetMeta())
	r.Use(middlewares.CheckCustomer(container.MContainer))

	v1 := r.Group("/api/v1")
	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON("pong")
	})

	v1.Get("/index", controller.Index().Get)

	productRouter := v1.Group("/product")
	productRouter.Get("/search", controller.Product().Search)
	productRouter.Get("/by_id/:id", controller.Product().ByID)
	productRouter.Get("/by_category_id/:category_id", controller.Product().ListByCategoryID)

	orderRouter := v1.Group("/order")
	orderRouter.Post("/", controller.Order().Create)

	fileRouter := v1.Group("/file")
	fileRouter.Get("/:filename", controller.File().Get)
	fileRouter.Post("/upload", controller.File().Upload)
	fileRouter.Post("/uploads", controller.File().Uploads)
}
