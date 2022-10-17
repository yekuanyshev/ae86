package repository

import (
	"context"
	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}

func (repo *OrderItemRepository) Create(ctx context.Context, orderItem model.OrderItem) (id int64, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"OrderItemRepository.Create failed",
				zap.Error(err),
				zap.Any("orderItem", orderItem),
			)
		}
	}()

	err = repo.db.Model(&model.OrderItem{}).Create(&orderItem).Error
	id = orderItem.ID
	return
}
