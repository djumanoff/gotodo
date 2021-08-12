package todo

import "github.com/djumanoff/gotodo/pkg/cqrses"

type CommandCreateTodo struct {
	Todo
}

func (cmd *CommandCreateTodo) Exec(svc interface{}) ([]cqrses.Event, interface{}, error) {
	todo, err := svc.(Service).NewTodo(cmd.Title, cmd.Body)
	return nil, todo, err
}
