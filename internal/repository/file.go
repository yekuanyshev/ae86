package repository

import (
	"context"
	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{db: db}
}

func (repo *FileRepository) ByName(ctx context.Context, name string) (result model.File, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"FileRepository.ByName failed",
				zap.Error(err),
				zap.String("name", name),
			)
		}
	}()

	err = repo.db.Model(&model.File{}).
		Where("name = ?", name).
		First(&result).Error
	return
}

func (repo *FileRepository) Create(ctx context.Context, file model.File) (id int64, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"FileRepository.Create failed",
				zap.Error(err),
				zap.Any("file", file),
			)
		}
	}()

	err = repo.db.Model(&model.File{}).Create(&file).Error
	id = file.ID
	return
}
