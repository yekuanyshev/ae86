package service

import "github.com/supernova0730/ae86/internal/interfaces/container"

type OrderService struct {
	mc container.IMainContainer
}

func NewOrderService(mc container.IMainContainer) *OrderService {
	return &OrderService{mc: mc}
}
