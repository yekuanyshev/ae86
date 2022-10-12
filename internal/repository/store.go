package repository

import (
	"context"
	"github.com/supernova0730/ae86/internal/repository/model"
	"github.com/supernova0730/ae86/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type StoreRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *StoreRepository {
	return &StoreRepository{db: db}
}

func (repo *StoreRepository) ByID(ctx context.Context, id int64) (result model.Store, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("StoreRepository.ByID failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.
		Model(&model.Store{}).
		Where("id = ?", id).
		First(&result).Error
	return
}
