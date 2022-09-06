package registry

import (
	"template/interface/controller"
	ip "template/interface/presenter"
	ir "template/interface/repository"
	"template/use_case/interactor"
	up "template/usecase/presenter"
	ur "template/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository, r.NewUserPresenter(), ir.NewDBRepository(r.db))
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserpresenter()
}
