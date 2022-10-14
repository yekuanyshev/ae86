package service

import (
	"context"
	"github.com/supernova0730/ae86/pkg/minio"
)

type IFileStorage interface {
	Upload(ctx context.Context, file *minio.File) (filename string, err error)
	Download(ctx context.Context, filename string) (file *minio.File, err error)
}
