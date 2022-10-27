package repository

import (
	"context"

	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) ActiveProductByIDAndStoreID(ctx context.Context, id int64, storeID int64) (result model.Product, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"ProductRepository.ActiveProductByIDAndStoreID failed",
				zap.Error(err),
				zap.Int64("id", id),
				zap.Int64("storeID", storeID),
			)
		}
	}()

	err = repo.db.Model(&model.Product{}).
		Joins("INNER JOIN category ON category.id = product.category_id").
		Where("category.store_id = ?", storeID).
		Where("product.id = ?", id).
		Where("category.is_active = ?", true).
		Where("product.is_active = ?", true).
		First(&result).Error
	return
}

func (repo *ProductRepository) ListActiveByCategoryIDAndStoreID(ctx context.Context, categoryID int64, storeID int64) (result []model.Product, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"ProductRepository.ListActiveByCategoryIDAndStoreID failed",
				zap.Error(err),
				zap.Int64("categoryID", categoryID),
				zap.Int64("storeID", storeID),
			)
		}
	}()

	err = repo.db.Model(&model.Product{}).
		Joins("INNER JOIN category ON category.id = product.category_id").
		Where("category.store_id = ?", storeID).
		Where("category.id = ?", categoryID).
		Where("category.is_active = ?", true).
		Where("product.is_active = ?", true).
		Find(&result).Error
	return
}

func (repo *ProductRepository) ListActiveByStoreIDAndSearchText(ctx context.Context, storeID int64, searchText string) (result []model.Product, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"ProductRepository.ListActiveByStoreIDAndSearchText",
				zap.Error(err),
				zap.Int64("storeID", storeID),
				zap.String("searchText", searchText),
			)
		}
	}()

	err = repo.db.Model(&model.Product{}).
		Joins("INNER JOIN category ON category.id = product.category_id").
		Where("category.store_id = ?", storeID).
		Where("category.is_active = ?", true).
		Where("product.is_active = ?", true).
		Where("LOWER(product.title) like LOWER(?)", searchText+"%").
		Find(&result).Error
	return
}
