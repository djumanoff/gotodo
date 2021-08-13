package todo

import (
	"github.com/djumanoff/gotodo/pkg/cqrses"
	"github.com/djumanoff/gotodo/pkg/utils"
)

type CommandCreateTodo struct {
	Todo
}

func (cmd *CommandCreateTodo) Exec(svc interface{}) ([]cqrses.Event, interface{}, error) {
	events := []cqrses.Event{}
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
	todo, err := svc.(Service).GetTodos()
	return nil, todo, err
}
