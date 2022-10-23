package app

import (
	"context"
	"github.com/supernova0730/ae86/config"
	"github.com/supernova0730/ae86/internal/connections"
	"github.com/supernova0730/ae86/internal/transport/rest"
	"github.com/supernova0730/ae86/pkg/logger"
	"github.com/supernova0730/ae86/pkg/minio"
	"github.com/supernova0730/ae86/pkg/postgres"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	err := config.Load("")
	if err != nil {
		logger.Log.Fatal("failed to load config", zap.Error(err))
	}

	logger.Log.Info("config loaded", zap.Any("config", config.Global))

	err = connections.DBConnect(postgres.Config{
		Host:     config.Global.DBHost,
		Port:     config.Global.DBPort,
		User:     config.Global.DBUser,
		Password: config.Global.DBPassword,
		Name:     config.Global.DBName,
		SSLMode:  config.Global.DBSSLMode,
	})
	if err != nil {
		logger.Log.Fatal("failed to connect to database", zap.Error(err))
	}

	logger.Log.Info("connected to database...")

	if config.Global.MinioEnabled {
		err = connections.MinioConnect(context.Background(), minio.Config{
			Host:     config.Global.MinioHost,
			Port:     config.Global.MinioPort,
			User:     config.Global.MinioUser,
			Password: config.Global.MinioPassword,
			UseSSL:   config.Global.MinioUseSSL,
			Bucket:   config.Global.MinioBucket,
		})
		if err != nil {
			logger.Log.Fatal("failed to connect to minio", zap.Error(err))
		}
		logger.Log.Info("connected to minio...")
	}

	server := rest.New(rest.Config{
		Host:      config.Global.HTTPHost,
		Port:      config.Global.HTTPPort,
		TLSEnable: config.Global.HTTPTLSEnable,
		CertFile:  config.Global.HTTPCertFile,
		KeyFile:   config.Global.HTTPKeyFile,
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
