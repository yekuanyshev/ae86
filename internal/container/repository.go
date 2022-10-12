package container

import (
	"github.com/supernova0730/ae86/internal/interfaces/repository"
	repository2 "github.com/supernova0730/ae86/internal/repository"
	"gorm.io/gorm"
)

type RepositoryContainer struct {
	banner    repository.IBannerRepository
	category  repository.ICategoryRepository
	customer  repository.ICustomerRepository
	manager   repository.IManagerRepository
	order     repository.IOrderRepository
	orderItem repository.IOrderItemRepository
	product   repository.IProductRepository
	store     repository.IStoreRepository
}

func NewRepositoryContainer(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		banner:    repository2.NewBannerRepository(db),
		category:  repository2.NewCategoryRepository(db),
		customer:  repository2.NewCustomerRepository(db),
		manager:   repository2.NewManagerRepository(db),
		order:     repository2.NewOrderRepository(db),
		orderItem: repository2.NewOrderItemRepository(db),
		product:   repository2.NewProductRepository(db),
		store:     repository2.NewStoreRepository(db),
	}
}

func (rc *RepositoryContainer) Banner() repository.IBannerRepository {
	return rc.banner
}

func (rc *RepositoryContainer) Category() repository.ICategoryRepository {
	return rc.category
}

func (rc *RepositoryContainer) Customer() repository.ICustomerRepository {
	return rc.customer
}

func (rc *RepositoryContainer) Manager() repository.IManagerRepository {
	return rc.manager
}

func (rc *RepositoryContainer) Order() repository.IOrderRepository {
	return rc.order
}

func (rc *RepositoryContainer) OrderItem() repository.IOrderItemRepository {
	return rc.orderItem
}

func (rc *RepositoryContainer) Product() repository.IProductRepository {
	return rc.product
}

func (rc *RepositoryContainer) Store() repository.IStoreRepository {
	return rc.store
}
