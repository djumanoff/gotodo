package todo

type Todo struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
	Body   string `json:"body"`
}

type TodoUpdate struct {
	ID     int64   `json:"-"`
	Title  *string `json:"title"`
	Status *string `json:"status"`
	Body   *string `json:"body"`
}
