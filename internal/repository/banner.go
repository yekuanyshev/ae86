package repository

import (
	"context"
	"github.com/supernova0730/ae86/internal/repository/model"
	"github.com/supernova0730/ae86/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BannerRepository struct {
	db *gorm.DB
}

func NewBannerRepository(db *gorm.DB) *BannerRepository {
	return &BannerRepository{db: db}
}

func (repo *BannerRepository) ListByStoreID(ctx context.Context, storeID int64) (result []model.Banner, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("BannerRepository.ListByStoreID failed", zap.Error(err), zap.Int64("storeID", storeID))
		}
	}()

	err = repo.db.
		Model(&model.Banner{}).
		Where("store_id = ?", storeID).
		Find(&result).Error
	return
}
