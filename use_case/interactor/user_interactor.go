package interactor

import (
	"errors"
	"template/domain/model"
	"template/use_case/presenter"
	"template/use_case/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
	DBRepostory    repository.DBRepository
}

type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
}

func (us *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}

func (us *userInteractor) Create(u *model.User) (*model.User, error) {
	data, err := us.DBRepostory.Transaction(func(i interface{}) (interface{}, error) {
		u, err := us.UserRepository.Create(u)

		// do other process here

		return u, err
	})

	user, ok := data.(*model.User)

	if !ok {
		return nil, errors.New("cast error")
	}

	if !errors.Is(err, nil) {
		return nil, err
	}

	return user, nil
}

func NewUserInteractor(
	r repository.UserRepository,
	p presenter.UserPresenter,
	d repository.DBRepository) UserInteractor {
	return &userInteractor{r, p, d}
}
