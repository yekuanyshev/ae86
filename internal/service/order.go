package service

import "github.com/supernova0730/ae86/internal/interfaces/container"

type OrderService struct {
	repository container.IRepository
}

func NewOrderService(repository container.IRepository) *OrderService {
	return &OrderService{repository: repository}
}
