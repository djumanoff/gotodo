package todo

import (
	"github.com/l00p8/l00p8"
)

type Repository interface {
	Update(upd *TodoUpdate) error

	Create(todo *Todo) (int64, error)

	FindAll(query *TodoQuery, params *l00p8.ListParams) ([]*Todo, error)

	FindById(id int64) (*Todo, error)

	Delete(id int64) error

	Health() error
}
