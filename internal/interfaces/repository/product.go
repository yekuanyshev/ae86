package repository

import (
	"context"
	"github.com/supernova0730/ae86/internal/repository/model"
)

type IProductRepository interface {
	ByID(ctx context.Context, id int64) (result model.Product, err error)
	ListActiveByCategoryID(ctx context.Context, categoryID int64) (result []model.Product, err error)
	ListActiveByStoreIDAndSearchText(ctx context.Context, storeID int64, searchText string) (result []model.Product, err error)
}
