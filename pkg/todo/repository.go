package todo

import http_helper "github.com/djumanoff/gotodo/pkg/http-helper"

type Repository interface {
	Update(upd *TodoUpdate) error

	Create(todo *Todo) error

	FindAll(params *http_helper.ListParams) ([]*Todo, error)

	FindById(id int64) (*Todo, error)

	Delete(id int64) error
}
