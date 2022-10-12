package service

import "github.com/supernova0730/ae86/internal/interfaces/container"

type CategoryService struct {
	repository container.IRepository
}

func NewCategoryService(repository container.IRepository) *CategoryService {
	return &CategoryService{repository: repository}
}
