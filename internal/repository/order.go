package repository

import (
	"context"

	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (repo *OrderRepository) ByID(ctx context.Context, id int64) (result model.Order, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"OrderRepository.ByID failed",
				zap.Error(err),
				zap.Int64("id", id),
			)
		}
	}()

	err = repo.db.Model(&model.Order{}).
		Where("id = ?", id).
		First(&result).Error
	return
}

func (repo *OrderRepository) Create(ctx context.Context, order model.Order) (id int64, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"OrderRepository.Create failed",
				zap.Error(err),
				zap.Any("order", order),
			)
		}
	}()

	err = repo.db.Model(&model.Order{}).Create(&order).Error
	id = order.ID
	return
}
