package service

import (
	"context"
	"errors"
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerService struct {
	mc container.IMainContainer
}

func NewCustomerService(mc container.IMainContainer) *CustomerService {
	return &CustomerService{mc: mc}
}

func (s *CustomerService) CheckExistence(ctx context.Context, customer model.Customer) (id int64, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"CustomerService.CheckExistence failed",
				zap.Error(err),
				zap.Any("customer", customer),
			)
		}
	}()

	oldCustomer, err := s.mc.Repositories().Customer().ByExternalID(ctx, customer.ExternalID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		id, err = s.mc.Repositories().Customer().Create(ctx, customer)
		return
	}

	err = s.mc.Repositories().Customer().UpdateByExternalID(ctx, oldCustomer.ExternalID, customer)
	if err != nil {
		return
	}

	return oldCustomer.ID, nil
}
