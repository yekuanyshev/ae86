package repository

import (
	"context"
	"github.com/supernova0730/ae86/internal/model"
)

type IBannerRepository interface {
	ListByStoreID(ctx context.Context, storeID int64) (result []model.Banner, err error)
}
