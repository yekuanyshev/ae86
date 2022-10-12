package service

import (
	"context"
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"github.com/supernova0730/ae86/internal/views"
)

type IndexService struct {
	service    container.IService
	repository container.IRepository
}

func NewIndexService(repository container.IRepository) *IndexService {
	return &IndexService{repository: repository}
}

func (service *IndexService) Get(ctx context.Context, storeID int64) (view views.Index, err error) {
	// todo: check customer existence

	store, err := service.repository.Store().ByID(ctx, storeID)
	if err != nil {
		return
	}

	banners, err := service.repository.Banner().ListByStoreID(ctx, store.ID)
	if err != nil {
		return
	}

	categories, err := service.repository.Category().ListActiveByStoreID(ctx, store.ID)
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
