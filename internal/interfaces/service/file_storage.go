package service

import (
	"context"
	"github.com/supernova0730/ae86/internal/model"
)

type IFileStorage interface {
	Upload(ctx context.Context, file model.File) (filename string, err error)
	Download(ctx context.Context, filename string) (file model.File, err error)
}
