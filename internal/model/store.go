package model

import "time"

type Store struct {
	ID        int64     `gorm:"column:id"`
	Title     string    `gorm:"column:title"`
	ManagerID int64     `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Store) TableName() string {
	return "store"
}
