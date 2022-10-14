package service

import "github.com/supernova0730/ae86/internal/interfaces/container"

type OrderItemService struct {
	mc container.IMainContainer
}

func NewOrderItemService(mc container.IMainContainer) *OrderItemService {
	return &OrderItemService{mc: mc}
}
