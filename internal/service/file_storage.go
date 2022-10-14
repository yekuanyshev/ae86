package service

import (
	"context"
	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/pkg/minio"
	"github.com/supernova0730/ae86/pkg/uuid"
	"go.uber.org/zap"
	"path/filepath"
)

type FileStorageService struct {
	client *minio.Client
}

func NewFileStorageService(client *minio.Client) *FileStorageService {
	return &FileStorageService{client: client}
}

func (s *FileStorageService) Upload(ctx context.Context, file *minio.File) (filename string, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"FileStorageService.Upload failed",
				zap.Error(err),
				zap.Any("file", file),
			)
		}
	}()

	filename = s.generateFilename(filepath.Ext(file.Name))
	file.Name = filename
	err = s.client.Upload(ctx, file)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (s *FileStorageService) Download(ctx context.Context, filename string) (file *minio.File, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"FileStorageService.Download failed",
				zap.Error(err),
				zap.String("filename", filename),
			)
		}
	}()

	return s.client.Download(ctx, filename)
}

func (s *FileStorageService) generateFilename(ext string) string {
	return uuid.Generate() + ext
}
