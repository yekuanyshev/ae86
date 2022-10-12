package service

import (
	"context"
	"errors"
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"github.com/supernova0730/ae86/internal/model"
	"gorm.io/gorm"
)

type CustomerService struct {
	repository container.IRepository
}

func NewCustomerService(repository container.IRepository) *CustomerService {
	return &CustomerService{repository: repository}
}

func (s *CustomerService) CheckExistence(ctx context.Context, customer model.Customer) (err error) {
	oldCustomer, err := s.repository.Customer().ByExternalID(ctx, customer.ExternalID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		_, err = s.repository.Customer().Create(ctx, customer)
		return
	}

	return s.repository.Customer().UpdateByExternalID(ctx, oldCustomer.ExternalID, customer)
}
