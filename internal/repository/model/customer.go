package model

import "time"

type Customer struct {
	ID         int64     `gorm:"column:id"`
	ExternalID int64     `gorm:"column:external_id"`
	Username   string    `gorm:"column:username"`
	Phone      string    `gorm:"column:phone"`
	FirstName  string    `gorm:"column:first_name"`
	LastName   string    `gorm:"column:last_name"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}

func (Customer) TableName() string {
	return "customer"
}
