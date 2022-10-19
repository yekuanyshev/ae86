package container

import "github.com/supernova0730/ae86/internal/interfaces/repository"

type IRepository interface {
	Banner() repository.IBannerRepository
	Category() repository.ICategoryRepository
	File() repository.IFileRepository
	Manager() repository.IManagerRepository
	Order() repository.IOrderRepository
	OrderItem() repository.IOrderItemRepository
	Product() repository.IProductRepository
	Store() repository.IStoreRepository
}
