package service

import "github.com/supernova0730/ae86/internal/interfaces/container"

type BannerService struct {
	mc container.IMainContainer
}

func NewBannerService(mc container.IMainContainer) *BannerService {
	return &BannerService{mc: mc}
}
