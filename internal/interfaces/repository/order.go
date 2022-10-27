package repository

import (
	"context"

	"github.com/supernova0730/ae86/internal/model"
)

type IOrderRepository interface {
	ByID(ctx context.Context, id int64) (result model.Order, err error)
	Create(ctx context.Context, order model.Order) (id int64, err error)
}
