package app

import (
	"github.com/supernova0730/ae86/config"
	"github.com/supernova0730/ae86/internal/container"
	"github.com/supernova0730/ae86/internal/transport"
	"github.com/supernova0730/ae86/internal/transport/rest"
	"github.com/supernova0730/ae86/pkg/logger"
	"github.com/supernova0730/ae86/pkg/postgres"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	conf, err := config.Load("")
	if err != nil {
		logger.Log.Fatal("failed to load config", zap.Error(err))
	}

	logger.Log.Info("config loaded", zap.Any("config", conf))

	db, err := postgres.Connect(postgres.Config{
		Host:     conf.DBHost,
		Port:     conf.DBPort,
		User:     conf.DBUser,
		Password: conf.DBPassword,
		Name:     conf.DBName,
		SSLMode:  conf.DBSSLMode,
	})
	if err != nil {
		logger.Log.Fatal("failed to connect to database", zap.Error(err))
	}

	logger.Log.Info("connected to database...")

	container.MContainer.Init(db)

	transport.Start(transport.Config{
		Rest: rest.Config{
			Host:      conf.HTTPHost,
			Port:      conf.HTTPPort,
			TLSEnable: conf.HTTPTLSEnable,
			CertFile:  conf.HTTPCertFile,
			KeyFile:   conf.HTTPKeyFile,
		},
	})

	logger.Log.Info("http server started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	logger.Log.Info("shutting down...")
}
