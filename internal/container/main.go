package container

import (
	"sync"

	"github.com/supernova0730/ae86/internal/interfaces/container"
)

var MContainer = &mainContainer{}

type mainContainer struct {
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
			mc.repositories = NewRepositoryContainer()
		}
	})
	return mc.repositories
}
