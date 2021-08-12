package todo

type Service interface {
	NewTodo(title string) (*Todo, error)
	Done(id int64) error
}

type svc struct {
	Repository
}
