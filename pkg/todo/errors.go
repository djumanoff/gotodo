package todo

import "errors"

var (
	ErrTodoNotFound    = errors.New("Todo not found.")
	ErrTodoCreation    = errors.New("Error occured while creating todo.")
	ErrNothingToUpdate = errors.New("Nothing to update.")
)
