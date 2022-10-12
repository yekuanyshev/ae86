package service

import "github.com/supernova0730/ae86/internal/interfaces/container"

type OrderItemService struct {
	repository container.IRepository
}

func NewOrderItemService(repository container.IRepository) *OrderItemService {
	return &OrderItemService{repository: repository}
}
