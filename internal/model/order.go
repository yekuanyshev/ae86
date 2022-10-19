package model

import "time"

type Order struct {
	ID                    int64     `gorm:"column:id"`
	CustomerID            string    `gorm:"column:customer_id"`
	StoreID               int64     `gorm:"column:store_id"`
	Address               string    `gorm:"column:address"`
	Phone                 string    `gorm:"column:phone"`
	Name                  string    `gorm:"column:name"`
	Comment               string    `gorm:"column:comment"`
	AllergiesInfo         string    `gorm:"column:allergies_info"`
	CancellationReason    string    `gorm:"column:cancellation_reason"`
	NeedKitchenAppliances bool      `gorm:"column:need_kitchen_appliances"`
	State                 string    `gorm:"column:state"`
	PaymentMethod         string    `gorm:"column:payment_state"`
	DeliveryMethod        string    `gorm:"column:delivery_method"`
	CreatedAt             time.Time `gorm:"column:created_at"`
	UpdatedAt             time.Time `gorm:"column:updated_at"`
}

func (Order) TableName() string {
	return "orders"
}
