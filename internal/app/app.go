package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/supernova0730/ae86/config"
	"github.com/supernova0730/ae86/internal/connections"
	"github.com/supernova0730/ae86/internal/transport/rest"
	"github.com/supernova0730/ae86/pkg/logger"
	"github.com/supernova0730/ae86/pkg/minio"
	"github.com/supernova0730/ae86/pkg/postgres"
	"go.uber.org/zap"
)

func Run() {
	err := config.Load("")
	if err != nil {
		logger.Log.Fatal("failed to load config", zap.Error(err))
	}

	logger.Log.Info("config loaded", zap.Any("config", config.Global))

	err = connections.DBConnect(postgres.Config{
		Host:     config.Global.DB.Host,
		Port:     config.Global.DB.Port,
		User:     config.Global.DB.User,
		Password: config.Global.DB.Password,
		Name:     config.Global.DB.Name,
		SSLMode:  config.Global.DB.SSLMode,
	})
	if err != nil {
		logger.Log.Fatal("failed to connect to database", zap.Error(err))
	}

	logger.Log.Info("connected to database...")

	if config.Global.Minio.Enabled {
		err = connections.MinioConnect(context.Background(), minio.Config{
			Host:     config.Global.Minio.Host,
			Port:     config.Global.Minio.Port,
			User:     config.Global.Minio.User,
			Password: config.Global.Minio.Password,
			UseSSL:   config.Global.Minio.UseSSL,
			Bucket:   config.Global.Minio.Bucket,
		})
		if err != nil {
			logger.Log.Fatal("failed to connect to minio", zap.Error(err))
		}
		logger.Log.Info("connected to minio...")
	}

	server := rest.New(rest.Config{
		Host:      config.Global.HTTP.Host,
		Port:      config.Global.HTTP.Port,
		TLSEnable: config.Global.HTTP.TLSEnable,
		CertFile:  config.Global.HTTP.CertFile,
		KeyFile:   config.Global.HTTP.KeyFile,
	})

	server.StartAsync()
	logger.Log.Info("http server started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	logger.Log.Info("shutting down...")

	if err = server.Shutdown(); err != nil {
		logger.Log.Error("failed to stop rest server", zap.Error(err))
	}
}
