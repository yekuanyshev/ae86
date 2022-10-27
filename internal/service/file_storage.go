package service

import (
	"bytes"
	"context"
	"io"
	"path/filepath"

	"github.com/supernova0730/ae86/internal/logger"
	"github.com/supernova0730/ae86/internal/model"
	"github.com/supernova0730/ae86/pkg/minio"
	"github.com/supernova0730/ae86/pkg/uuid"
	"go.uber.org/zap"
)

type FileStorageService struct {
	client *minio.Client
}

func NewFileStorageService(client *minio.Client) *FileStorageService {
	return &FileStorageService{client: client}
}

func (s *FileStorageService) Upload(ctx context.Context, file model.File) (filename string, err error) {
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
	err = s.client.Upload(ctx, &minio.Object{
		Content:     bytes.NewBuffer(file.Content),
		Name:        filename,
		Size:        file.Size,
		ContentType: file.ContentType,
	})
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (s *FileStorageService) Download(ctx context.Context, filename string) (file model.File, err error) {
	defer func() {
		if err != nil {
			logger.LogWithCtx(ctx).Error(
				"FileStorageService.Download failed",
				zap.Error(err),
				zap.String("filename", filename),
			)
		}
	}()

	minioObject, err := s.client.Download(ctx, filename)
	if err != nil {
		return
	}

	objectContent, err := io.ReadAll(minioObject.Content)
	if err != nil {
		return
	}

	return model.File{
		Name:        minioObject.Name,
		ContentType: minioObject.ContentType,
		Size:        minioObject.Size,
		Content:     objectContent,
	}, nil
}

func (s *FileStorageService) generateFilename(ext string) string {
	return uuid.Generate() + ext
}
