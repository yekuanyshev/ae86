package container

import (
	irepository "github.com/supernova0730/ae86/internal/interfaces/repository"
	"github.com/supernova0730/ae86/internal/repository"
	"gorm.io/gorm"
	"sync"
)

type repositoryContainer struct {
	db *gorm.DB

	bannerInit sync.Once
	banner     irepository.IBannerRepository

	categoryInit sync.Once
	category     irepository.ICategoryRepository

	fileInit sync.Once
	file     irepository.IFileRepository

	managerInit sync.Once
	manager     irepository.IManagerRepository

	orderInit sync.Once
	order     irepository.IOrderRepository

	orderItemInit sync.Once
	orderItem     irepository.IOrderItemRepository

	productInit sync.Once
	product     irepository.IProductRepository

	storeInit sync.Once
	store     irepository.IStoreRepository
}

func NewRepositoryContainer(db *gorm.DB) *repositoryContainer {
	return &repositoryContainer{db: db}
}

func (rc *repositoryContainer) Banner() irepository.IBannerRepository {
	rc.bannerInit.Do(func() {
		if rc.banner == nil {
			rc.banner = repository.NewBannerRepository(rc.db)
		}
	})
	return rc.banner
}

func (rc *repositoryContainer) Category() irepository.ICategoryRepository {
	rc.categoryInit.Do(func() {
		if rc.category == nil {
			rc.category = repository.NewCategoryRepository(rc.db)
		}
	})
	return rc.category
}

func (rc *repositoryContainer) File() irepository.IFileRepository {
	rc.fileInit.Do(func() {
		if rc.file == nil {
			rc.file = repository.NewFileRepository(rc.db)
		}
	})
	return rc.file
}

func (rc *repositoryContainer) Manager() irepository.IManagerRepository {
	rc.managerInit.Do(func() {
		if rc.manager == nil {
			rc.manager = repository.NewManagerRepository(rc.db)
		}
	})
	return rc.manager
}

func (rc *repositoryContainer) Order() irepository.IOrderRepository {
	rc.orderInit.Do(func() {
		if rc.order == nil {
			rc.order = repository.NewOrderRepository(rc.db)
		}
	})
	return rc.order
}

func (rc *repositoryContainer) OrderItem() irepository.IOrderItemRepository {
	rc.orderItemInit.Do(func() {
		if rc.orderItem == nil {
			rc.orderItem = repository.NewOrderItemRepository(rc.db)
		}
	})
	return rc.orderItem
}

func (rc *repositoryContainer) Product() irepository.IProductRepository {
	rc.productInit.Do(func() {
		if rc.product == nil {
			rc.product = repository.NewProductRepository(rc.db)
		}
	})
	return rc.product
}

func (rc *repositoryContainer) Store() irepository.IStoreRepository {
	rc.storeInit.Do(func() {
		if rc.store == nil {
			rc.store = repository.NewStoreRepository(rc.db)
		}
	})
	return rc.store
}
