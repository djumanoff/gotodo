package todo

type Repository interface {
	FindById(id int64) (*Todo, error)
}
