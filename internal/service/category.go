package service

import "github.com/supernova0730/ae86/internal/interfaces/container"

type CategoryService struct {
	mc container.IMainContainer
}

func NewCategoryService(mc container.IMainContainer) *CategoryService {
	return &CategoryService{mc: mc}
}
