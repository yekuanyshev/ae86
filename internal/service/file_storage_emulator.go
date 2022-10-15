package service

import (
	"context"
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/internal/model"
	"github.com/supernova0730/ae86/pkg/uuid"
	"go.uber.org/zap"
	"path/filepath"
)

type FileStorageEmulatorService struct {
	mc container.IMainContainer
}

func NewFileStorageEmulatorService(mc container.IMainContainer) *FileStorageEmulatorService {
	return &FileStorageEmulatorService{mc: mc}
}

func (s *FileStorageEmulatorService) Upload(ctx context.Context, file model.File) (filename string, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"FileStorageEmulatorService.Upload failed",
				zap.Error(err),
				zap.Any("file", file),
			)
		}
	}()

	file.Name = s.generateFilename(filepath.Ext(file.Name))
	_, err = s.mc.Repositories().File().Create(ctx, file)
	if err != nil {
		return
	}

	return file.Name, nil
}

func (s *FileStorageEmulatorService) Download(ctx context.Context, filename string) (file model.File, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"FileStorageEmulatorService.Download failed",
				zap.Error(err),
				zap.String("filename", filename),
			)
		}
	}()

	return s.mc.Repositories().File().ByName(ctx, filename)
}

func (s *FileStorageEmulatorService) generateFilename(ext string) string {
	return uuid.Generate() + ext
}
