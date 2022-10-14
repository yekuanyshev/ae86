package service

import "github.com/supernova0730/ae86/internal/interfaces/container"

type StoreService struct {
	mc container.IMainContainer
}

func NewStoreService(mc container.IMainContainer) *StoreService {
	return &StoreService{mc: mc}
}
