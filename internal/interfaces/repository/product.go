package repository

import (
	"context"

	"github.com/supernova0730/ae86/internal/model"
)

type IProductRepository interface {
	ActiveProductByIDAndStoreID(ctx context.Context, id int64, storeID int64) (result model.Product, err error)
	ListActiveByCategoryIDAndStoreID(ctx context.Context, categoryID int64, storeID int64) (result []model.Product, err error)
	ListActiveByStoreIDAndSearchText(ctx context.Context, storeID int64, searchText string) (result []model.Product, err error)
}
