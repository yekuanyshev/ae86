package container

import (
	"github.com/supernova0730/ae86/internal/controllers"
	"github.com/supernova0730/ae86/internal/interfaces/container"
)

type ControllerContainer struct {
	banner    *controllers.BannerController
	category  *controllers.CategoryController
	customer  *controllers.CustomerController
	index     *controllers.IndexController
	manager   *controllers.ManagerController
	order     *controllers.OrderController
	orderItem *controllers.OrderItemController
	product   *controllers.ProductController
	store     *controllers.StoreController
}

func NewControllerContainer(service container.IService) *ControllerContainer {
	return &ControllerContainer{
		banner:    controllers.NewBannerController(service),
		category:  controllers.NewCategoryController(service),
		customer:  controllers.NewCustomerController(service),
		index:     controllers.NewIndexController(service),
		manager:   controllers.NewManagerController(service),
		order:     controllers.NewOrderController(service),
		orderItem: controllers.NewOrderItemController(service),
		product:   controllers.NewProductController(service),
		store:     controllers.NewStoreController(service),
	}
}

func (cc *ControllerContainer) Banner() *controllers.BannerController {
	return cc.banner
}

func (cc *ControllerContainer) Category() *controllers.CategoryController {
	return cc.category
}

func (cc *ControllerContainer) Customer() *controllers.CustomerController {
	return cc.customer
}

func (cc *ControllerContainer) Index() *controllers.IndexController {
	return cc.index
}

func (cc *ControllerContainer) Manager() *controllers.ManagerController {
	return cc.manager
}

func (cc *ControllerContainer) Order() *controllers.OrderController {
	return cc.order
}

func (cc *ControllerContainer) OrderItem() *controllers.OrderItemController {
	return cc.orderItem
}

func (cc *ControllerContainer) Product() *controllers.ProductController {
	return cc.product
}

func (cc *ControllerContainer) Store() *controllers.StoreController {
	return cc.store
}
