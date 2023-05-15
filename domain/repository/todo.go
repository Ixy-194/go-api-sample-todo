package repository

import "app/domain/model"

type Todo interface {
	Create(t *model.Todo) error
	Delete(id int) error
	Update(t *model.Todo) error
	Find(id int) (*model.Todo, error)
	FindAll() ([]*model.Todo, error)
}
