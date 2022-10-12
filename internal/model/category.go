package model

import "time"

type Category struct {
	ID        int64     `gorm:"column:id"`
	Title     string    `gorm:"column:title"`
	Image     string    `gorm:"column:image"`
	StoreID   int64     `gorm:"column:store_id"`
	IsActive  bool      `gorm:"column:is_active"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Category) TableName() string {
	return "category"
}
