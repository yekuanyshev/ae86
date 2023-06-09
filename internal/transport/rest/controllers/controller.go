package controllers

import (
	"github.com/supernova0730/ae86/internal/interfaces/container"
)

type controllerContainer struct {
	banner   *BannerController
	category *CategoryController
	file     *FileController
	index    *IndexController
	manager  *ManagerController
	order    *OrderController
	product  *ProductController
	store    *StoreController
}

func NewContainer(service container.IService) *controllerContainer {
	return &controllerContainer{
		banner:   NewBannerController(service),
		category: NewCategoryController(service),
		file:     NewFileController(service),
		index:    NewIndexController(service),
		manager:  NewManagerController(service),
		order:    NewOrderController(service),
		product:  NewProductController(service),
		store:    NewStoreController(service),
	}
}

func (cc *controllerContainer) Banner() *BannerController {
	return cc.banner
}

func (cc *controllerContainer) Category() *CategoryController {
	return cc.category
}

func (cc *controllerContainer) File() *FileController {
	return cc.file
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

func (cc *controllerContainer) Product() *ProductController {
	return cc.product
}

func (cc *controllerContainer) Store() *StoreController {
	return cc.store
}
