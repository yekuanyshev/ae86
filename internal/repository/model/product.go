package model

import "time"

type Product struct {
	ID          int64     `gorm:"column:id"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	Price       int       `gorm:"column:price"`
	Image       string    `gorm:"column:image"`
	IsActive    bool      `gorm:"column:is_active"`
	CategoryID  int64     `gorm:"column:category_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Product) TableName() string {
	return "product"
}
