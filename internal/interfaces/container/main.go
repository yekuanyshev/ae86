package container

type IMainContainer interface {
	Services() IService
	Repositories() IRepository
}
