package application

import (
	"hexagonal-architecture-sample/server/application/model"
)

type UserInterface interface {
	Create(user model.User) error
	GetAll() ([]model.User, error)
}

type User struct {
	Interface UserInterface
}

func (u *User) Create(user model.User) error {
	return u.Interface.Create(user)
}

func (u *User) GetAll() ([]model.User, error) {
	return u.Interface.GetAll()
}
