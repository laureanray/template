package repository

import "template/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
}
