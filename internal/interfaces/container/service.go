package container

import "github.com/supernova0730/ae86/internal/interfaces/service"

type IService interface {
	Banner() service.IBannerService
	Category() service.ICategoryService
	Customer() service.ICustomerService
	Index() service.IIndexService
	Manager() service.IManagerService
	Order() service.IOrderService
	OrderItem() service.IOrderItemService
	Product() service.IProductService
	Store() service.IStoreService
}
