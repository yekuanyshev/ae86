package repository

import (
	"context"

	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IngredientRepository struct {
	db *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) *IngredientRepository {
	return &IngredientRepository{db: db}
}

func (repo *IngredientRepository) ListActiveIngredientsByProductID(ctx context.Context, productID int64) (result []model.Ingredient, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"IngredientRepository.ListActiveIngredientsByProductID failed",
				zap.Error(err),
				zap.Int64("productID", productID),
			)
		}
	}()

	err = repo.db.Model(&model.Ingredient{}).
		Where("product_id = ?", productID).
		Where("is_active = ?", true).
		Find(&result).Error
	return
}
