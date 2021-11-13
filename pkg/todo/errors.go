package todo

import "errors"

var (
	ErrTodoNotFound    = errors.New("todo not found")
	ErrTodoCreation    = errors.New("error occured while creating todo")
	ErrNothingToUpdate = errors.New("nothing to update")
)
