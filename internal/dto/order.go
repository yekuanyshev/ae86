package dto

type OrderCreateDTO struct {
	Name                  string               `json:"name"`
	Phone                 string               `json:"phone"`
	Address               string               `json:"address"`
	Comment               string               `json:"comment"`
	AllergiesInfo         string               `json:"allergiesInfo"`
	CancellationReason    string               `json:"cancellationReason"`
	NeedKitchenAppliances bool                 `json:"needKitchenAppliances"`
	PaymentMethod         string               `json:"paymentMethod"`  // CASH, CARD
	DeliveryMethod        string               `json:"deliveryMethod"` // BY_COURIER, TAKEAWAY
	Items                 []OrderItemCreateDTO `json:"items"`
}

type OrderItemCreateDTO struct {
	ProductID int64 `json:"productID"`
	Amount    int   `json:"amount"`
}
