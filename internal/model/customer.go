package model

import "time"

type Customer struct {
	ID         int64     `gorm:"column:id"`
	ExternalID string    `gorm:"column:external_id"`
	Username   string    `gorm:"column:username"`
	FirstName  string    `gorm:"column:first_name"`
	LastName   string    `gorm:"column:last_name"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}

func (Customer) TableName() string {
	return "customer"
}
