package enums

type (
	PaymentMethod  string
	OrderState     string
	DeliveryMethod string
)

const (
	PaymentMethodCash = "CASH"
	PaymentMethodCard = "CARD"

	OrderStatePending            = "PENDING"
	OrderStateCanceled           = "CANCELED"
	OrderStateAccepted           = "ACCEPTED"
	OrderStateDeliveryInProgress = "DELIVERY_IN_PROGRESS"
	OrderStateDelivered          = "DELIVERED"

	DeliveryMethodByCourier = "BY_COURIER"
	DeliveryMethodTakeaway  = "TAKEAWAY"
)
