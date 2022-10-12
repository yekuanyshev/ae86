package service

import (
	"context"
	"github.com/supernova0730/ae86/internal/views"
)

type IProductService interface {
	ListByCategoryID(ctx context.Context, categoryID int64) (result []views.ProductShort, err error)
	Search(ctx context.Context, storeID int64, searchText string) (view views.ProductSearchResult, err error)
}
