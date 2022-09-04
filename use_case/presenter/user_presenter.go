package presenter

import "template/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
