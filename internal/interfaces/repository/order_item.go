package repository

import (
	"context"

	"github.com/supernova0730/ae86/internal/model"
)

type IOrderItemRepository interface {
	Create(ctx context.Context, orderItem model.OrderItem) (id int64, err error)
}
