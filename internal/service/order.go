package service

import (
	"context"
	"errors"
	"github.com/supernova0730/ae86/internal/dto"
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/internal/model"
	"go.uber.org/zap"
)

type OrderService struct {
	mc container.IMainContainer
}

func NewOrderService(mc container.IMainContainer) *OrderService {
	return &OrderService{mc: mc}
}

func (s *OrderService) Create(ctx context.Context, customerID int64, storeID int64, orderCreateDTO dto.OrderCreateDTO) (id int64, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"OrderService.Create failed",
				zap.Error(err),
				zap.Any("orderCreateDTO", orderCreateDTO),
			)
		}
	}()

	if len(orderCreateDTO.Items) == 0 {
		err = errors.New("order items is empty")
		return
	}

	orderID, err := s.mc.Repositories().Order().Create(ctx, model.Order{
		CustomerID:            customerID,
		StoreID:               storeID,
		Address:               orderCreateDTO.Address,
		Comment:               orderCreateDTO.Comment,
		AllergiesInfo:         orderCreateDTO.AllergiesInfo,
		CancellationReason:    orderCreateDTO.CancellationReason,
		NeedKitchenAppliances: orderCreateDTO.NeedKitchenAppliances,
		State:                 "PENDING",
		PaymentMethod:         orderCreateDTO.PaymentMethod,
		DeliveryMethod:        orderCreateDTO.DeliveryMethod,
	})
	if err != nil {
		return
	}

	for _, orderItemDTO := range orderCreateDTO.Items {
		_, err = s.mc.Repositories().OrderItem().Create(ctx, model.OrderItem{
			OrderID:   orderID,
			ProductID: orderItemDTO.ProductID,
			Amount:    orderItemDTO.Amount,
		})
		if err != nil {
			return
		}
	}

	return orderID, nil
}
