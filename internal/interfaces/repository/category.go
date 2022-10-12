package repository

import (
	"context"
	"github.com/supernova0730/ae86/internal/repository/model"
)

type ICategoryRepository interface {
	ListActiveByStoreID(ctx context.Context, storeID int64) (result []model.Category, err error)
}
