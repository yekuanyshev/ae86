package service

import "github.com/supernova0730/ae86/internal/interfaces/container"

type CustomerService struct {
	repository container.IRepository
}

func NewCustomerService(repository container.IRepository) *CustomerService {
	return &CustomerService{repository: repository}
}
