package rest

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/internal/container"
	"github.com/supernova0730/ae86/pkg/logger"
	"go.uber.org/zap"
)

func Start(config Config, controller *container.ControllerContainer) {
	app := fiber.New(fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})

	RegisterMiddlewares(app)
	RegisterRoutes(app, controller)

	address := config.Address()

	var err error
	if config.TLSEnable {
		err = app.ListenTLS(address, config.CertFile, config.KeyFile)
	} else {
		err = app.Listen(address)
	}
	if err != nil {
		logger.Log.Fatal("failed to start rest server", zap.Error(err))
	}
}

func RegisterMiddlewares(app *fiber.App) {
	app.Use(SetContextHolder())
	app.Use(SetStoreID())
}
