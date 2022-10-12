package container

import (
	"github.com/supernova0730/ae86/internal/interfaces/container"
	controllers2 "github.com/supernova0730/ae86/internal/transport/rest/controllers"
)

type ControllerContainer struct {
	banner    *controllers2.BannerController
	category  *controllers2.CategoryController
	customer  *controllers2.CustomerController
	index     *controllers2.IndexController
	manager   *controllers2.ManagerController
	order     *controllers2.OrderController
	orderItem *controllers2.OrderItemController
	product   *controllers2.ProductController
	store     *controllers2.StoreController
}

func NewControllerContainer(service container.IService) *ControllerContainer {
	return &ControllerContainer{
		banner:    controllers2.NewBannerController(service),
		category:  controllers2.NewCategoryController(service),
		customer:  controllers2.NewCustomerController(service),
		index:     controllers2.NewIndexController(service),
		manager:   controllers2.NewManagerController(service),
		order:     controllers2.NewOrderController(service),
		orderItem: controllers2.NewOrderItemController(service),
		product:   controllers2.NewProductController(service),
		store:     controllers2.NewStoreController(service),
	}
}

func (cc *ControllerContainer) Banner() *controllers2.BannerController {
	return cc.banner
}

func (cc *ControllerContainer) Category() *controllers2.CategoryController {
	return cc.category
}

func (cc *ControllerContainer) Customer() *controllers2.CustomerController {
	return cc.customer
}

func (cc *ControllerContainer) Index() *controllers2.IndexController {
	return cc.index
}

func (cc *ControllerContainer) Manager() *controllers2.ManagerController {
	return cc.manager
}

func (cc *ControllerContainer) Order() *controllers2.OrderController {
	return cc.order
}

func (cc *ControllerContainer) OrderItem() *controllers2.OrderItemController {
	return cc.orderItem
}

func (cc *ControllerContainer) Product() *controllers2.ProductController {
	return cc.product
}

func (cc *ControllerContainer) Store() *controllers2.StoreController {
	return cc.store
}
