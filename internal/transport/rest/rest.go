package rest

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/pkg/logger"
	"go.uber.org/zap"
	"net"
)

type Config struct {
	Host      string
	Port      string
	TLSEnable bool
	CertFile  string
	KeyFile   string
}

func (c Config) Address() string {
	return net.JoinHostPort(c.Host, c.Port)
}

func Start(config Config) {
	app := fiber.New(fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})

	RegisterRoutes(app)

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
