package service

import "github.com/supernova0730/ae86/internal/interfaces/container"

type StoreService struct {
	repository container.IRepository
}

func NewStoreService(repository container.IRepository) *StoreService {
	return &StoreService{repository: repository}
}
