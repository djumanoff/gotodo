package todo

import (
	"github.com/djumanoff/gotodo/pkg/utils"
)

type Repository interface {
	Update(upd *TodoUpdate) error

	Create(todo *Todo) (int64, error)

	FindAll(query *TodoQuery, params *utils.ListParams) ([]*Todo, error)

	FindById(id int64) (*Todo, error)

	Delete(id int64) error

	Health() error
}
