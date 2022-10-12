package repository

import (
	"context"
	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (repo *CustomerRepository) ByExternalID(ctx context.Context, externalID string) (result model.Customer, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"CustomerRepository.ByExternalID failed",
				zap.Error(err),
				zap.String("externalID", externalID),
			)
		}
	}()

	err = repo.db.Model(&model.Customer{}).
		Where("external_id = ?", externalID).
		First(&result).Error
	return
}

func (repo *CustomerRepository) Create(ctx context.Context, customer model.Customer) (id int64, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"CustomerRepository.Create failed",
				zap.Error(err),
				zap.Any("customer", customer),
			)
		}
	}()

	err = repo.db.Model(&model.Customer{}).Create(&customer).Error
	id = customer.ID
	return
}

func (repo *CustomerRepository) UpdateByExternalID(ctx context.Context, externalID string, customer model.Customer) (err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"CustomerRepository.UpdateByExternalID failed",
				zap.Error(err),
				zap.String("externalID", externalID),
				zap.Any("customer", customer),
			)
		}
	}()

	err = repo.db.Model(&model.Customer{}).
		Where("external_id = ?", externalID).
		Updates(&customer).Error
	return
}
