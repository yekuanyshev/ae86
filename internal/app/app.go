package app

import (
	"context"
	"github.com/supernova0730/ae86/config"
	"github.com/supernova0730/ae86/internal/container"
	"github.com/supernova0730/ae86/internal/transport"
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
	ctx := context.Background()

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

	minioConfig := minio.Config{
		Host:     conf.MinioHost,
		Port:     conf.MinioPort,
		User:     conf.MinioUser,
		Password: conf.MinioPassword,
		UseSSL:   conf.MinioUseSSL,
		Bucket:   conf.MinioBucket,
	}

	minioClient, err := minio.Connect(ctx, minioConfig)
	if err != nil {
		logger.Log.Fatal("failed to connect to file storage", zap.Error(err))
	}

	logger.Log.Info("connected to file storage...")

	container.MContainer.Init(db, minioClient)

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
