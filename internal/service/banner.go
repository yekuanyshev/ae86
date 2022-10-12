package service

import "github.com/supernova0730/ae86/internal/interfaces/container"

type BannerService struct {
	repository container.IRepository
}

func NewBannerService(repository container.IRepository) *BannerService {
	return &BannerService{repository: repository}
}
