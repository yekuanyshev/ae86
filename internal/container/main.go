package container

import (
	"github.com/supernova0730/ae86/internal/interfaces/container"
	"gorm.io/gorm"
	"sync"
)

var MContainer = &mainContainer{}

type mainContainer struct {
	db *gorm.DB

	servicesInit sync.Once
	services     container.IService

	repositoriesInit sync.Once
	repositories     container.IRepository
}

func (mc *mainContainer) Services() container.IService {
	mc.servicesInit.Do(func() {
		if mc.services == nil {
			mc.services = NewServiceContainer()
		}
	})
	return mc.services
}

func (mc *mainContainer) Repositories() container.IRepository {
	mc.repositoriesInit.Do(func() {
		if mc.repositories == nil {
			mc.repositories = NewRepositoryContainer(mc.db)
		}
	})
	return mc.repositories
}

func (mc *mainContainer) Init(db *gorm.DB) {
	mc.db = db
}
