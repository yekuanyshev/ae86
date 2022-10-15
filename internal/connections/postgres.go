package connections

import (
	"github.com/supernova0730/ae86/pkg/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func DBConnect(config postgres.Config) error {
	db, err := postgres.Connect(config)
	if err != nil {
		return err
	}
	DBConn = db
	return nil
}
