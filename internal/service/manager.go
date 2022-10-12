package service

import "github.com/supernova0730/ae86/internal/interfaces/container"

type ManagerService struct {
	repository container.IRepository
}

func NewManagerService(repository container.IRepository) *ManagerService {
	return &ManagerService{repository: repository}
}
