package repository

import (
	"context"
	"github.com/supernova0730/ae86/internal/repository/model"
	"github.com/supernova0730/ae86/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) ByID(ctx context.Context, id int64) (result model.Product, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ProductRepository.ByID failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.
		Model(&model.Product{}).
		Where("id = ?", id).
		First(&result).Error
	return
}

func (repo *ProductRepository) ListActiveByCategoryID(ctx context.Context, categoryID int64) (result []model.Product, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ProductRepository.ListActiveByCategoryID failed", zap.Error(err), zap.Int64("categoryID", categoryID))
		}
	}()

	err = repo.db.
		Model(&model.Product{}).
		Where("category_id = ?", categoryID).
		Where("is_active = ?", true).
		Find(&result).Error
	return
}

func (repo *ProductRepository) ListActiveByStoreIDAndSearchText(ctx context.Context, storeID int64, searchText string) (result []model.Product, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ProductRepository.ListActiveByStoreIDAndSearchText", zap.Error(err), zap.Int64("storeID", storeID), zap.String("searchText", searchText))
		}
	}()

	err = repo.db.
		Model(&model.Product{}).
		Joins("INNER JOIN category ON category.id = product.category_id").
		Where("category.store_id = ?", storeID).
		Where("product.is_active = ?", true).
		Where("LOWER(product.title) like LOWER(?)", searchText+"%").
		Find(&result).Error
	return
}
