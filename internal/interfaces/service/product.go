package service

import (
	"context"
	"github.com/supernova0730/ae86/internal/views"
)

type IProductService interface {
	ByID(ctx context.Context, storeID int64, id int64) (view views.ProductFull, err error)
	ListByCategoryID(ctx context.Context, storeID int64, categoryID int64) (result []views.ProductShort, err error)
	Search(ctx context.Context, storeID int64, searchText string) (view views.ProductSearchResult, err error)
}
