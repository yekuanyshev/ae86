package repository

import (
	"context"

	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) ListActiveByStoreID(ctx context.Context, storeID int64) (result []model.Category, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"CategoryRepository.ListByStoreID failed",
				zap.Error(err),
				zap.Int64("storeID", storeID),
			)
		}
	}()

	err = repo.db.Model(&model.Category{}).
		Where("store_id = ?", storeID).
		Where("is_active = ?", true).
		Find(&result).Error
	return
}
