package container

import (
	irepository "github.com/supernova0730/ae86/internal/interfaces/repository"
	"github.com/supernova0730/ae86/internal/repository"
	"gorm.io/gorm"
)

type RepositoryContainer struct {
	banner    irepository.IBannerRepository
	category  irepository.ICategoryRepository
	customer  irepository.ICustomerRepository
	manager   irepository.IManagerRepository
	order     irepository.IOrderRepository
	orderItem irepository.IOrderItemRepository
	product   irepository.IProductRepository
	store     irepository.IStoreRepository
}

func NewRepositoryContainer(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		banner:    repository.NewBannerRepository(db),
		category:  repository.NewCategoryRepository(db),
		customer:  repository.NewCustomerRepository(db),
		manager:   repository.NewManagerRepository(db),
		order:     repository.NewOrderRepository(db),
		orderItem: repository.NewOrderItemRepository(db),
		product:   repository.NewProductRepository(db),
		store:     repository.NewStoreRepository(db),
	}
}

func (rc *RepositoryContainer) Banner() irepository.IBannerRepository {
	return rc.banner
}

func (rc *RepositoryContainer) Category() irepository.ICategoryRepository {
	return rc.category
}

func (rc *RepositoryContainer) Customer() irepository.ICustomerRepository {
	return rc.customer
}

func (rc *RepositoryContainer) Manager() irepository.IManagerRepository {
	return rc.manager
}

func (rc *RepositoryContainer) Order() irepository.IOrderRepository {
	return rc.order
}

func (rc *RepositoryContainer) OrderItem() irepository.IOrderItemRepository {
	return rc.orderItem
}

func (rc *RepositoryContainer) Product() irepository.IProductRepository {
	return rc.product
}

func (rc *RepositoryContainer) Store() irepository.IStoreRepository {
	return rc.store
}
