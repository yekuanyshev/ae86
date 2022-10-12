package postgres

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)
import "gorm.io/driver/postgres"

func Connect(config Config) (*gorm.DB, error) {
	dialector := postgres.Open(config.DSN())

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			Colorful: true,
			LogLevel: logger.Info,
		},
	)

	return gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
	})
}
