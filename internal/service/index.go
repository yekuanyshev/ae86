package service

import (
	"context"
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"github.com/supernova0730/ae86/internal/views"
)

type IndexService struct {
	mc container.IMainContainer
}

func NewIndexService(mc container.IMainContainer) *IndexService {
	return &IndexService{mc: mc}
}

func (service *IndexService) Get(ctx context.Context, storeID int64) (view views.Index, err error) {
	store, err := service.mc.Repositories().Store().ByID(ctx, storeID)
	if err != nil {
		return
	}

	banners, err := service.mc.Repositories().Banner().ListByStoreID(ctx, store.ID)
	if err != nil {
		return
	}

	categories, err := service.mc.Repositories().Category().ListActiveByStoreID(ctx, store.ID)
	if err != nil {
		return
	}

	view.StoreTitle = store.Title
	for _, banner := range banners {
		view.Banners = append(view.Banners, banner.Image)
	}

	for _, category := range categories {
		view.Categories = append(view.Categories, views.Category{
			ID:    category.ID,
			Title: category.Title,
			Image: category.Image,
		})
	}
	return
}
