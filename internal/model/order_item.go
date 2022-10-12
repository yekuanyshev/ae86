package model

type OrderItem struct {
	ID        int64 `gorm:"column:id"`
	OrderID   int64 `gorm:"column:order_id"`
	ProductID int64 `gorm:"column:product_id"`
	Amount    int   `gorm:"column:amount"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
