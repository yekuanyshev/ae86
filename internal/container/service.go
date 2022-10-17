package container

import (
	"github.com/supernova0730/ae86/config"
	"github.com/supernova0730/ae86/internal/connections"
	iservice "github.com/supernova0730/ae86/internal/interfaces/service"
	"github.com/supernova0730/ae86/internal/service"
	"sync"
)

type serviceContainer struct {
	bannerInit sync.Once
	banner     iservice.IBannerService

	categoryInit sync.Once
	category     iservice.ICategoryService

	customerInit sync.Once
	customer     iservice.ICustomerService

	fileStorageInit sync.Once
	fileStorage     iservice.IFileStorage

	indexInit sync.Once
	index     iservice.IIndexService

	managerInit sync.Once
	manager     iservice.IManagerService

	orderInit sync.Once
	order     iservice.IOrderService

	productInit sync.Once
	product     iservice.IProductService

	storeInit sync.Once
	store     iservice.IStoreService
}

func NewServiceContainer() *serviceContainer {
	return &serviceContainer{}
}

func (sc *serviceContainer) Banner() iservice.IBannerService {
	sc.bannerInit.Do(func() {
		if sc.banner == nil {
			sc.banner = service.NewBannerService(MContainer)
		}
	})
	return sc.banner
}

func (sc *serviceContainer) Category() iservice.ICategoryService {
	sc.categoryInit.Do(func() {
		if sc.category == nil {
			sc.category = service.NewCategoryService(MContainer)
		}
	})
	return sc.category
}

func (sc *serviceContainer) Customer() iservice.ICustomerService {
	sc.customerInit.Do(func() {
		if sc.customer == nil {
			sc.customer = service.NewCustomerService(MContainer)
		}
	})
	return sc.customer
}

func (sc *serviceContainer) FileStorage() iservice.IFileStorage {
	sc.fileStorageInit.Do(func() {
		if sc.fileStorage == nil {
			if config.Global.MinioEnabled {
				sc.fileStorage = service.NewFileStorageService(connections.MinioConn)
			} else {
				sc.fileStorage = service.NewFileStorageEmulatorService(MContainer)
			}
		}
	})
	return sc.fileStorage
}

func (sc *serviceContainer) Index() iservice.IIndexService {
	sc.indexInit.Do(func() {
		if sc.index == nil {
			sc.index = service.NewIndexService(MContainer)
		}
	})
	return sc.index
}

func (sc *serviceContainer) Manager() iservice.IManagerService {
	sc.managerInit.Do(func() {
		if sc.manager == nil {
			sc.manager = service.NewManagerService(MContainer)
		}
	})
	return sc.manager
}

func (sc *serviceContainer) Order() iservice.IOrderService {
	sc.orderInit.Do(func() {
		if sc.order == nil {
			sc.order = service.NewOrderService(MContainer)
		}
	})
	return sc.order
}

func (sc *serviceContainer) Product() iservice.IProductService {
	sc.productInit.Do(func() {
		if sc.product == nil {
			sc.product = service.NewProductService(MContainer)
		}
	})
	return sc.product
}

func (sc *serviceContainer) Store() iservice.IStoreService {
	sc.storeInit.Do(func() {
		if sc.store == nil {
			sc.store = service.NewStoreService(MContainer)
		}
	})
	return sc.store
}
