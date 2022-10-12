package model

import "time"

type Manager struct {
	ID        int64     `gorm:"column:id"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (Manager) TableName() string {
	return "manager"
}
