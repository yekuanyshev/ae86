package model

type Ingredient struct {
	ID        int64  `gorm:"column:id"`
	Title     string `gorm:"column:title"`
	Price     int    `gorm:"column:price"`
	IsActive  bool   `gorm:"column:is_active"`
	ProductID int64  `gorm:"column:product_id"`
}

func (Ingredient) TableName() string {
	return "ingredient"
}
