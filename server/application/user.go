package application

import (
	"hexagonal-architecture-sample/server/application/model"
	"hexagonal-architecture-sample/server/application/repository"
)

type User struct {
	Interface repository.User
}

func (u *User) Create(user model.User) error {
	return u.Interface.Create(user)
}

func (u *User) GetAll() ([]model.User, error) {
	return u.Interface.GetAll()
}

func (u *User) Update(user model.User) error {
	return u.Interface.Update(user)
}

func (u *User) GetByID(id string) (*model.User, error) {
	return u.Interface.GetByID(id)
}
