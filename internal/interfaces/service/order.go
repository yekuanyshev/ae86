package service

import (
	"context"
	"github.com/supernova0730/ae86/internal/dto"
)

type IOrderService interface {
	Create(ctx context.Context, customerID int64, storeID int64, orderCreateDTO dto.OrderCreateDTO) (id int64, err error)
}
