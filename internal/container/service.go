package container

import (
	"github.com/supernova0730/ae86/internal/interfaces/container"
	iservice "github.com/supernova0730/ae86/internal/interfaces/service"
	"github.com/supernova0730/ae86/internal/service"
)

type ServiceContainer struct {
	banner    iservice.IBannerService
	category  iservice.ICategoryService
	customer  iservice.ICustomerService
	index     iservice.IIndexService
	manager   iservice.IManagerService
	order     iservice.IOrderService
	orderItem iservice.IOrderItemService
	product   iservice.IProductService
	store     iservice.IStoreService
}

func NewServiceContainer(repository container.IRepository) *ServiceContainer {
	return &ServiceContainer{
		banner:    service.NewBannerService(repository),
		category:  service.NewCategoryService(repository),
		customer:  service.NewCustomerService(repository),
		index:     service.NewIndexService(repository),
		manager:   service.NewManagerService(repository),
		order:     service.NewOrderService(repository),
		orderItem: service.NewOrderItemService(repository),
		product:   service.NewProductService(repository),
		store:     service.NewStoreService(repository),
	}
}

func (sc ServiceContainer) Banner() iservice.IBannerService {
	return sc.banner
}

func (sc ServiceContainer) Category() iservice.ICategoryService {
	return sc.category
}

func (sc ServiceContainer) Customer() iservice.ICustomerService {
	return sc.customer
}

func (sc ServiceContainer) Index() iservice.IIndexService {
	return sc.index
}

func (sc ServiceContainer) Manager() iservice.IManagerService {
	return sc.manager
}

func (sc ServiceContainer) Order() iservice.IOrderService {
	return sc.order
}

func (sc ServiceContainer) OrderItem() iservice.IOrderItemService {
	return sc.orderItem
}

func (sc ServiceContainer) Product() iservice.IProductService {
	return sc.product
}

func (sc ServiceContainer) Store() iservice.IStoreService {
	return sc.store
}
