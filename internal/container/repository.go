package container

import (
	"sync"

	"github.com/supernova0730/ae86/internal/connections"
	irepository "github.com/supernova0730/ae86/internal/interfaces/repository"
	"github.com/supernova0730/ae86/internal/repository"
)

type repositoryContainer struct {
	bannerInit sync.Once
	banner     irepository.IBannerRepository

	categoryInit sync.Once
	category     irepository.ICategoryRepository

	fileInit sync.Once
	file     irepository.IFileRepository

	ingredientInit sync.Once
	ingredient     irepository.IIngredientRepository

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

func NewRepositoryContainer() *repositoryContainer {
	return &repositoryContainer{}
}

func (rc *repositoryContainer) Banner() irepository.IBannerRepository {
	rc.bannerInit.Do(func() {
		if rc.banner == nil {
			rc.banner = repository.NewBannerRepository(connections.DBConn)
		}
	})
	return rc.banner
}

func (rc *repositoryContainer) Category() irepository.ICategoryRepository {
	rc.categoryInit.Do(func() {
		if rc.category == nil {
			rc.category = repository.NewCategoryRepository(connections.DBConn)
		}
	})
	return rc.category
}

func (rc *repositoryContainer) File() irepository.IFileRepository {
	rc.fileInit.Do(func() {
		if rc.file == nil {
			rc.file = repository.NewFileRepository(connections.DBConn)
		}
	})
	return rc.file
}

func (rc *repositoryContainer) Ingredient() irepository.IIngredientRepository {
	rc.ingredientInit.Do(func() {
		if rc.ingredient == nil {
			rc.ingredient = repository.NewIngredientRepository(connections.DBConn)
		}
	})
	return rc.ingredient
}

func (rc *repositoryContainer) Manager() irepository.IManagerRepository {
	rc.managerInit.Do(func() {
		if rc.manager == nil {
			rc.manager = repository.NewManagerRepository(connections.DBConn)
		}
	})
	return rc.manager
}

func (rc *repositoryContainer) Order() irepository.IOrderRepository {
	rc.orderInit.Do(func() {
		if rc.order == nil {
			rc.order = repository.NewOrderRepository(connections.DBConn)
		}
	})
	return rc.order
}

func (rc *repositoryContainer) OrderItem() irepository.IOrderItemRepository {
	rc.orderItemInit.Do(func() {
		if rc.orderItem == nil {
			rc.orderItem = repository.NewOrderItemRepository(connections.DBConn)
		}
	})
	return rc.orderItem
}

func (rc *repositoryContainer) Product() irepository.IProductRepository {
	rc.productInit.Do(func() {
		if rc.product == nil {
			rc.product = repository.NewProductRepository(connections.DBConn)
		}
	})
	return rc.product
}

func (rc *repositoryContainer) Store() irepository.IStoreRepository {
	rc.storeInit.Do(func() {
		if rc.store == nil {
			rc.store = repository.NewStoreRepository(connections.DBConn)
		}
	})
	return rc.store
}
