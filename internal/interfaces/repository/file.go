package repository

import (
	"context"
	"github.com/supernova0730/ae86/internal/model"
)

type IFileRepository interface {
	ByName(ctx context.Context, name string) (result model.File, err error)
	Create(ctx context.Context, file model.File) (id int64, err error)
}
