package rest

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/supernova0730/ae86/pkg/logger"
	"go.uber.org/zap"
)

type Server struct {
	config Config
	app    *fiber.App
}

func New(config Config) *Server {
	app := fiber.New(fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})

	s := &Server{
		config: config,
		app:    app,
	}
	s.registerRoutes()
	return s
}

func (s *Server) Start() error {
	addr := s.config.Address()
	if s.config.TLSEnable {
		return s.app.ListenTLS(addr, s.config.CertFile, s.config.KeyFile)
	}
	return s.app.Listen(addr)
}

func (s *Server) StartAsync() {
	go func() {
		err := s.Start()
		if err != nil {
			logger.Log.Fatal("failed to start server", zap.Error(err))
		}
	}()
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
