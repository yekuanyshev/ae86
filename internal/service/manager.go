package service

import "github.com/supernova0730/ae86/internal/interfaces/container"

type ManagerService struct {
	mc container.IMainContainer
}

func NewManagerService(mc container.IMainContainer) *ManagerService {
	return &ManagerService{mc: mc}
}
