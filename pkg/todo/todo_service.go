package todo

type Service interface {
	NewTodo(title string, body string) (*Todo, error)

	Done(id int64) error
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
