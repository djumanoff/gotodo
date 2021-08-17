package todo

import "github.com/djumanoff/gotodo/pkg/utils"

type Service interface {
	NewTodo(title string, body string) (*Todo, error)

	Done(id int64) error

	GetTodos(query *TodoQuery, params *utils.ListParams) ([]*Todo, error)

	GetTodo(id int64) (*Todo, error)
}

func NewService(repo Repository) Service {
	return &defaultSvc{repo}
}

type defaultSvc struct {
	repo Repository
}

func (svc *defaultSvc) NewTodo(title string, body string) (*Todo, error) {
	todo := &Todo{}
	todo.Title = title
	todo.Body = body
	todo.Status = "todo"
	id, err := svc.repo.Create(todo)
	if err != nil {
		return nil, err
	}
	todo.ID = id
	return todo, nil
}

func (svc *defaultSvc) Done(id int64) error {
	status := "done"
	upd := &TodoUpdate{}
	upd.ID = id
	upd.Status = &status
	err := svc.repo.Update(upd)
	if err != nil {
		return err
	}
	return nil
}

func (svc *defaultSvc) GetTodos(query *TodoQuery, params *utils.ListParams) ([]*Todo, error) {
	return svc.repo.FindAll(query, params)
}

func (svc *defaultSvc) GetTodo(id int64) (*Todo, error) {
	return svc.repo.FindById(id)
}
