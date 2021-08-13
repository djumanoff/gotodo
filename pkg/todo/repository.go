package todo

import (
	http_helper "github.com/djumanoff/gotodo/pkg/utils"
)

type Repository interface {
	Update(upd *TodoUpdate) error

	Create(todo *Todo) (int64, error)

	FindAll(params *http_helper.ListParams) ([]*Todo, error)

	FindById(id int64) (*Todo, error)

	Delete(id int64) error

	Health() error
}
