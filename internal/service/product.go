package service

import (
	"context"
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"github.com/supernova0730/ae86/internal/views"
)

type ProductService struct {
	mc container.IMainContainer
}

func NewProductService(mc container.IMainContainer) *ProductService {
	return &ProductService{mc: mc}
}

func (service *ProductService) ByID(ctx context.Context, storeID int64, id int64) (view views.ProductFull, err error) {
	product, err := service.mc.Repositories().Product().ActiveProductByIDAndStoreID(ctx, id, storeID)
	if err != nil {
		return
	}

	return views.ProductFull{
		ID:          product.ID,
		Title:       product.Title,
		Price:       product.Price,
		Description: product.Description,
		Image:       product.Image,
	}, nil
}

func (service *ProductService) ListByCategoryID(ctx context.Context, storeID int64, categoryID int64) (result []views.ProductShort, err error) {
	products, err := service.mc.Repositories().Product().ListActiveByCategoryIDAndStoreID(ctx, categoryID, storeID)
	if err != nil {
		return
	}

	for _, product := range products {
		result = append(result, views.ProductShort{
			ID:    product.ID,
			Title: product.Title,
			Image: product.Image,
			Price: product.Price,
		})
	}
	return
}

func (service *ProductService) Search(ctx context.Context, storeID int64, searchText string) (view views.ProductSearchResult, err error) {
	products, err := service.mc.Repositories().Product().ListActiveByStoreIDAndSearchText(ctx, storeID, searchText)
	if err != nil {
		return
	}

	for _, product := range products {
		view.Products = append(view.Products, views.ProductShort{
			ID:    product.ID,
			Title: product.Title,
			Image: product.Image,
			Price: product.Price,
		})
	}
	return
}
