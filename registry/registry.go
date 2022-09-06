package registry

import (
	"template/interface/controller"

	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *regitry) NewAppController() controller.AppController {
	return controller.AppController{
		User: r.NewUserController(),
	}
}
