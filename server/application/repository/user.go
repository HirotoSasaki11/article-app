package repository

import "hexagonal-architecture-sample/server/application/model"

type User interface {
	Create(user model.User) error
	GetAll() ([]model.User, error)
	Update(user model.User) error
	GetByID(id string) (*model.User, error)
}
