package container

import "github.com/supernova0730/ae86/internal/interfaces/service"

type IService interface {
	Banner() service.IBannerService
	Category() service.ICategoryService
	Customer() service.ICustomerService
	FileStorage() service.IFileStorage
	Index() service.IIndexService
	Manager() service.IManagerService
	Order() service.IOrderService
	Product() service.IProductService
	Store() service.IStoreService
}
