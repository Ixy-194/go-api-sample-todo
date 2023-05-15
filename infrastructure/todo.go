package infrastructure

import (
	"app/domain/model"
	"app/domain/repository"

	"gorm.io/gorm"
)

type Todo struct {
	db *gorm.DB
}

func NewTodo(db *gorm.DB) repository.Todo {
	return &Todo{
		db: db,
	}
}

func (td *Todo) Create(t *model.Todo) error {
	if err := td.db.Create(t).Error; err != nil {
		return err
	}
	return nil
}

func (td *Todo) Update(t *model.Todo) error {
	if err := td.db.Save(t).Error; err != nil {
		return err
	}
	return nil
}

func (td *Todo) Delete(id int) error {
	if err := td.db.Where("id = ?", id).Delete(&model.Todo{}).Error; err != nil {
		return err
	}
	return nil
}

func (td *Todo) Find(id int) (*model.Todo, error) {
	var todo *model.Todo
	err := td.db.Where("id = ?", id).Take(&todo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return todo, nil
}
func (td *Todo) FindAll() ([]*model.Todo, error) {
	var todos []*model.Todo
	err := td.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}
