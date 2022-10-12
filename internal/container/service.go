package container

import (
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"github.com/supernova0730/ae86/internal/interfaces/service"
	service2 "github.com/supernova0730/ae86/internal/service"
)

type ServiceContainer struct {
	banner    service.IBannerService
	category  service.ICategoryService
	customer  service.ICustomerService
	index     service.IIndexService
	manager   service.IManagerService
	order     service.IOrderService
	orderItem service.IOrderItemService
	product   service.IProductService
	store     service.IStoreService
}

func NewServiceContainer(repository container.IRepository) *ServiceContainer {
	return &ServiceContainer{
		banner:    service2.NewBannerService(repository),
		category:  service2.NewCategoryService(repository),
		customer:  service2.NewCustomerService(repository),
		index:     service2.NewIndexService(repository),
		manager:   service2.NewManagerService(repository),
		order:     service2.NewOrderService(repository),
		orderItem: service2.NewOrderItemService(repository),
		product:   service2.NewProductService(repository),
		store:     service2.NewStoreService(repository),
	}
}

func (sc ServiceContainer) Banner() service.IBannerService {
	return sc.banner
}

func (sc ServiceContainer) Category() service.ICategoryService {
	return sc.category
}

func (sc ServiceContainer) Customer() service.ICustomerService {
	return sc.customer
}

func (sc ServiceContainer) Index() service.IIndexService {
	return sc.index
}

func (sc ServiceContainer) Manager() service.IManagerService {
	return sc.manager
}

func (sc ServiceContainer) Order() service.IOrderService {
	return sc.order
}

func (sc ServiceContainer) OrderItem() service.IOrderItemService {
	return sc.orderItem
}

func (sc ServiceContainer) Product() service.IProductService {
	return sc.product
}

func (sc ServiceContainer) Store() service.IStoreService {
	return sc.store
}
