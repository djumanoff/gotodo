package todo

import (
	"github.com/l00p8/cqrses"
	"github.com/l00p8/utils"
)

type CommandCreateTodo struct {
	Todo
}

func (cmd *CommandCreateTodo) Exec(svc interface{}) ([]cqrses.Event, interface{}, error) {
	var events []cqrses.Event
	todo, err := svc.(Service).NewTodo(cmd.Title, cmd.Body)
	if err != nil {
		return nil, nil, err
	}
	events = append(events, cqrses.NewEventWithJson("TodoCreated", todo))
	return events, todo, err
}

type CommandGetTodos struct {
	Query      *TodoQuery
	ListParams *utils.ListParams
}

func (cmd *CommandGetTodos) Exec(svc interface{}) ([]cqrses.Event, interface{}, error) {
	todos, err := svc.(Service).GetTodos(cmd.Query, cmd.ListParams)
	return nil, todos, err
}

type CommandGetTodoByID struct {
	ID int64
}

func (cmd *CommandGetTodoByID) Exec(svc interface{}) ([]cqrses.Event, interface{}, error) {
	todo, err := svc.(Service).GetTodo(cmd.ID)
	return nil, todo, err
}
