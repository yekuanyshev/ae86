package controllers

import (
	"github.com/supernova0730/ae86/internal/interfaces/container"
)

type controllerContainer struct {
	banner    *BannerController
	category  *CategoryController
	customer  *CustomerController
	index     *IndexController
	manager   *ManagerController
	order     *OrderController
	orderItem *OrderItemController
	product   *ProductController
	store     *StoreController
}

func NewContainer(service container.IService) *controllerContainer {
	return &controllerContainer{
		banner:    NewBannerController(service),
		category:  NewCategoryController(service),
		customer:  NewCustomerController(service),
		index:     NewIndexController(service),
		manager:   NewManagerController(service),
		order:     NewOrderController(service),
		orderItem: NewOrderItemController(service),
		product:   NewProductController(service),
		store:     NewStoreController(service),
	}
}

func (cc *controllerContainer) Banner() *BannerController {
	return cc.banner
}

func (cc *controllerContainer) Category() *CategoryController {
	return cc.category
}

func (cc *controllerContainer) Customer() *CustomerController {
	return cc.customer
}

func (cc *controllerContainer) Index() *IndexController {
	return cc.index
}

func (cc *controllerContainer) Manager() *ManagerController {
	return cc.manager
}

func (cc *controllerContainer) Order() *OrderController {
	return cc.order
}

func (cc *controllerContainer) OrderItem() *OrderItemController {
	return cc.orderItem
}

func (cc *controllerContainer) Product() *ProductController {
	return cc.product
}

func (cc *controllerContainer) Store() *StoreController {
	return cc.store
}
